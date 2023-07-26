package search

import (
  "fmt"
  "net/http"
  "net/http/httputil"
)

func SerachMessages() {
  // ref: https://qiita.com/taizo/items/c397dbfed7215969b0a5
  // url := "https://slack.com/api/search.messages?query=test&pretty=1"
  // url := "https://slack.com/api/search.messages?query=AI&pretty=1"
  url := "https://slack.com/api/search.messages?query=test&pretty=1"
  req, _ := http.NewRequest("GET", url, nil)
  req.Header.Set("Authorization", "Bearer xoxp-5634909965604-5632355661763-5656267501088-94f85e051307ec75eb82e0aa694666a5")

  dump, _ := httputil.DumpRequestOut(req, true)
  fmt.Printf("%s", dump)

  client := new(http.Client)
  resp, err := client.Do(req)
  fmt.Printf("%s", err)

  dumpResp, _ := httputil.DumpResponse(resp, true)
  fmt.Printf("%s", dumpResp)
}
