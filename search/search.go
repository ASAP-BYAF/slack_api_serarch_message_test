package search

import (
	"strings"
	"fmt"
    "net/http"
    "net/http/httputil"
)

func SerachMessages(query string) {

    // query 文字列の空白を %20 に置き換える。
	// query 文字列に空白が含まれているとエラーになる。
	query  = strings.ReplaceAll(query, " ", "%20")

	// slack api のエンドポイント
	url := "https://slack.com/api/search.messages?pretty=1"
	url += "&query=" + query
	
	// ヘッダーに app の token を追加。
	// ref: https://qiita.com/taizo/items/c397dbfed7215969b0a5
	req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer xoxp-5634909965604-5632355661763-5638582518147-6f843e20c76044a649bee7fd00b4d82c")

    dump, _ := httputil.DumpRequestOut(req, true)
    fmt.Printf("%s", dump)

    client := new(http.Client)
    resp, err := client.Do(req)
    fmt.Printf("%s", err)

    dumpResp, _ := httputil.DumpResponse(resp, true)
    fmt.Printf("%s", dumpResp)
}
