package alchemyapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

type stubHTTP struct {
	GetCount    int
	GetByte     []byte
	GetEndPoint string
	GetParams   url.Values

	PostCount    int
	PostByte     []byte
	PostEndPoint string
	PostParams   url.Values
}

func (h *stubHTTP) get(endPoint string, params url.Values, config *config) ([]byte, error) {
	h.GetCount++
	h.GetEndPoint = endPoint
	h.GetParams = params
	return h.GetByte, nil
}

func (h *stubHTTP) post(endPoint string, params url.Values, config *config) ([]byte, error) {
	h.PostCount++
	h.PostEndPoint = endPoint
	h.PostParams = params
	return h.PostByte, nil
}

func TestURLGetRankedImageKeywords(t *testing.T) {
	Convey("correct", t, func() {
		correctJSON := `
			{
			   "status": "REQUEST_STATUS",
			   "url": "REQUESTED_URL",
			   "totalTransactions": "42",
			   "imageKeywords": [{
			      "text": "DETECTED_KEYWORD",
                  "score": "4.2",
                  "knowledgeGraph":
                  {
                    "typeHierarchy": "DETECTED_TYPE_HIERARCHY"
                  }
               }]
            }`

		testToken := "testToken"
		client := New(testToken)

		Convey("URLGetRankedImageKeywords", func() {
			params := url.Values{}
			params.Add("url", "http://example.com")

			stub := &stubHTTP{}
			stub.PostByte = []byte(correctJSON)
			client.connection = stub

			res, err := client.URLGetRankedImageKeywords("http://example.com", false, false)
			So(err, ShouldBeNil)
			So(stub.PostParams, ShouldResemble, params)

			So(res.Status, ShouldEqual, "REQUEST_STATUS")
			So(res.URL, ShouldEqual, "REQUESTED_URL")
			So(res.TotalTransactions, ShouldEqual, "42")
			So(len(res.ImageKeywords), ShouldEqual, 1)

			keyword := res.ImageKeywords[0]
			So(keyword.Text, ShouldEqual, "DETECTED_KEYWORD")
			So(keyword.Score, ShouldEqual, "4.2")

		})
	})
}
