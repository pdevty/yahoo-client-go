yahoo-client-go
==================

[![GoDoc](https://godoc.org/github.com/pdevty/yahoo-client-go?status.svg)](https://godoc.org/github.com/pdevty/yahoo-client-go) 
[![Coverage Status](https://coveralls.io/repos/github/pdevty/yahoo-client-go/badge.svg?branch=master)](https://coveralls.io/github/pdevty/yahoo-client-go?branch=master)
[![Build Status](https://travis-ci.org/pdevty/yahoo-client-go.svg?branch=master)](https://travis-ci.org/pdevty/yahoo-client-go)

yahoo-client-go is a go client library for [yahoo japan web api](http://developer.yahoo.co.jp/sitemap/).


## Installation

```go
go get -u github.com/pdevty/yahoo-client-go
```

## Usage

```go
import yahoo "github.com/pdevty/yahoo-client-go"
```

```go
client := yahoo.NewClient("<put your appid>")

resultset, _ := client.ItemSearch(&yahoo.ItemSearchParam{
	Query: "vaio",
})

for _, v := range resultset.Result.Hit {
	fmt.Println(v.Category.Current.ID)
}
```

```go
res, err := client.ItemSearch(&yahoo.ItemSearchParam{})
res, err := client.CategoryRanking(&yahoo.CategoryRankingParam{})
res, err := client.CategorySearch(&yahoo.CategorySearchParam{})
res, err := client.ItemLookup(&yahoo.ItemLookupParam{})
res, err := client.QueryRanking(&yahoo.QueryRankingParam{})
res, err := client.GetModule(&yahoo.GetModuleParam{})
res, err := client.ShopCampaignSearch(&yahoo.ShopCampaignSearchParam{})
res, err := client.ReviewSearch(&yahoo.ReviewSearchParam{})
```

## Contributing

1. Fork ([https://github.com/pdevty/yahoo-client-go/fork](https://github.com/pdevty/yahoo-client-go/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request
