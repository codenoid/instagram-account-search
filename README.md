<img align="right" width="100" height="100" src="http://pn-marisa.go.id/tkn/uploads/2020/07/instagram-png-instagram-png-logo-1455.png">

# IG Account Search

Search instagram account, that's it, This libra ry intended for education purpose only

[![Go Report Card](https://goreportcard.com/badge/github.com/codenoid/instagram-account-search)](https://goreportcard.com/report/github.com/codenoid/instagram-account-search)
[![CodeFactor](https://www.codefactor.io/repository/github/codenoid/instagram-account-search/badge/master)](https://www.codefactor.io/repository/github/codenoid/instagram-account-search/overview/master)

## as HTTP/API Server

Hit this API to use

### Installation
```sh
go get github.com/codenoid/instagram-account-search/cmd/ig-account-search-server
ig-account-search-server -bind :3000
```

### Usage

```sh
curl 'http://localhost:3000/search?q=lola%20zieta'
```

## as Library

`igaccountsearch` package library

### Installation

```sh
go get -u github.com/codenoid/instagram-account-search
```

### Usage

```go
package main

import (
	"log"

	igaccountsearch "github.com/codenoid/instagram-account-search"
)

func main() {
    searchResult, err := igaccountsearch.UserSearch("lola zieta")
    log.Println(err)
    log.Println(searchResult) // igaccountsearch.IGSearchResult structs
}
```

## Credits

All the credits belongs to [@rakd](https://github.com/rakd/gin_sample/tree/6f6d31d29a81f4fcc7f59dd24399b0e5404cc2ed/app/libs/igsearch)

## Legal

This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Instagram or any of its affiliates or subsidiaries. This is an independent and unofficial software. Use at your own risk.