package search

import (
	"strings"
    "os"
	"fmt"
    "net/http"
    "net/http/httputil"
)

func SerachMessages(query string) {

    // slack api のエンドポイント
	url := "https://slack.com/api/search.messages?pretty=1"

    // クエリパラメータにクエリ文字列を追加。
	// 空白が含まれているとエラーになるので %20 に置き換える。
	query  = strings.ReplaceAll(query, " ", "%20")
	url += "&query=" + query
	
    // ヘッダーに app の認証トークンを追加。
	// ref: https://qiita.com/taizo/items/c397dbfed7215969b0a5
	req, _ := http.NewRequest("GET", url, nil)
	token := os.Getenv("SEARCH_APP_TOKEN")
    req.Header.Set("Authorization", "Bearer "+token)

    dump, _ := httputil.DumpRequestOut(req, true)
    fmt.Printf("%s", dump)

    client := new(http.Client)
    resp, err := client.Do(req)
    fmt.Printf("%s", err)

    dumpResp, _ := httputil.DumpResponse(resp, true)
    fmt.Printf("%s", dumpResp)
}
