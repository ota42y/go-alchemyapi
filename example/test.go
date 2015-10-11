package main

import (
	"bytes"
	"fmt"
	alchemyapi "github.com/ota42y/go-alchemyapi"
	"io/ioutil"
	"os"
)

func main() {
	token := os.Getenv("ALCHEMYAPI_TOKEN")
	if token == "" {
		fmt.Println("skip this test because no token")
		return
	}

	client := alchemyapi.New(token)
	res, err := client.URLGetRankedImageKeywords("https://pbs.twimg.com/profile_images/509356702265667584/_j6Y7hlU_400x400.png", true, true)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

	b, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		fmt.Println(err)
	} else {
		res, err = client.ImageGetRankedImageKeywords(bytes.NewReader(b), true, true)
		if err == nil {
			fmt.Println(res)
		} else {
			fmt.Println(err)
		}
	}
}
