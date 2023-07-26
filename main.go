package main

import (
  "fmt"

  "github.com/ASAP-BYAF/slack_api_serarch_message_test/search"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env")
	
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
	
  var query string = "open notion key"
  search.SerachMessages(query)
}
