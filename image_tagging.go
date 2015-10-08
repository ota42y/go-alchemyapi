package alchemyapi

import (
	"encoding/json"
	"net/url"
)

// URLGetRankedImageKeywordsResponse is URLGetRankedImageKeywords response struct
type URLGetRankedImageKeywordsResponse struct {
	Status            string
	URL               string
	KnowledgeGraph    string
	StatusInfo        string
	ImageKeywords     []ImageKeyword
	TotalTransactions string
}

// ImageKeyword is URLGetRankedImageKeywords response struct
type ImageKeyword struct {
	Text           string
	Score          string
	KnowledgeGraph KnowledgeGraph
}

// KnowledgeGraph is URLGetRankedImageKeywords response struct
type KnowledgeGraph struct {
	TypeHierarchy string
}

// URLGetRankedImageKeywords return  http://gateway-a.watsonplatform.net/calls/url/URLGetRankedImageKeywords response
func (api *AlchemyAPI) URLGetRankedImageKeywords(imageURL string, forceShowAll bool, knowledgeGraph bool) (URLGetRankedImageKeywordsResponse, error) {
	params := url.Values{}
	params.Add("url", imageURL)

	if forceShowAll {
		params.Add("forceShowAll", "0")
	}

	if knowledgeGraph {
		params.Add("knowledgeGraph", "0")
	}

	var res URLGetRankedImageKeywordsResponse

	b, err := api.connection.post("calls/url/URLGetRankedImageKeywords", params, api.c)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(b, &res)
	return res, err
}
