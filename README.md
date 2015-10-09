# go-alchemyapi
[AlchemyAPI](http://www.alchemyapi.com/) client for golang.

# Install
```bash
$ go get github.com/ota42y/go-alchemyapi
```

# Usage
```go
// need export ALCHEMYAPI_TOKEN=token

package main

import (
    "os"
    "fmt"
    "bytes"
    "io/ioutil"

    alchemyapi "github.com/ota42y/go-alchemyapi"
)

func main() {
    token := os.Getenv("ALCHEMYAPI_TOKEN")
    if token == "" {
        fmt.Println("skip this test because no token")
        return
    }

    client := alchemyapi.New(token)
    res, err := client.URLGetRankedImageKeywords("https://pbs.twimg.com/profile_images/509356702265667584/_j6Y7hlU_400x400.png", false, false)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}

```

# feature
- support image tagging api
