# curl2struct

`curl2struct` is a tool designed to convert curl commands into equivalent Golang structs. 
This can be particularly useful when working on Golang projects and needing to translate HTTP requests from curl commands into Go code.

## Installation

To install `curl2struct`, you can use `go get`:

```bash
go get -u github.com/yourusername/curl2struct
```

## Usage

```go
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
  -d '{"a": "b"}'
  --compressed`

func main() {
    c, err := curl2struct.NewCurl(curl)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(c)
    /*
        URL: https://github.com/3n0ugh/curl2struct/tree-commit-info/main
        Method: POST
        Headers: map[Accept:[application/json] Accept-Language:[en-US,en;q=0.9] Authority:[github.com] Content-Type:[application/json]]
        Body: {"a": "b"}
    */
}
```

* The NewCurl(curl string) function returns the CURL model.

```go
type CURL struct {
    URL     string
    Method  string
    Headers http.Header
    Body    []byte
}
```
