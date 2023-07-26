package search

import (
  "fmt"
  "net/http"
  "net/http/httputil"
)

func SerachMessages(query string) {
	url := "https://slack.com/api/search.messages?pretty=1"
	url += "&query=" + query
	
	// ref: https://qiita.com/taizo/items/c397dbfed7215969b0a5
	req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer xoxp-5634909965604-5632355661763-5638667613250-fac0549c90aba9a5ff5e39a6d553c43f")

    dump, _ := httputil.DumpRequestOut(req, true)
    fmt.Printf("%s", dump)

    client := new(http.Client)
    resp, err := client.Do(req)
    fmt.Printf("%s", err)

    dumpResp, _ := httputil.DumpResponse(resp, true)
    fmt.Printf("%s", dumpResp)
}
