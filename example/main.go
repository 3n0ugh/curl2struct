package main

import (
	"fmt"
	"github.com/3n0ugh/curl2struct"
	"log"
)

var curl = `curl 'https://github.com/3n0ugh/curl2struct/tree-commit-info/main' \
  -H 'authority: github.com' \
  -H 'accept: application/json' \
  -H 'accept-language: en-US,en;q=0.9' \
  -H 'content-type: application/json' \
  -H 'dnt: 1' \
  -H 'github-verified-fetch: true' \
  -H 'if-none-match: W/"ffd65006d577a49fe021f4af6116fa24"' \
  -H 'referer: https://github.com/3n0ugh/curl2struct' \
  -H 'sec-ch-ua: "Not_A Brand";v="8", "Chromium";v="120"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'x-requested-with: XMLHttpRequest' \
  --compressed`

func main() {
	c, err := curl2struct.NewCurl(curl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
}
