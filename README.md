# yahoo-client-go
==================

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
client = yahoo.NewClient("<Put your API key>")

resultset, err := client.ItemSearch(&ItemSearchParam{
    Query: "vaio"
})

```

## Contributing

1. Fork ([https://github.com/pdevty/yahoo-client-go/fork](https://github.com/pdevty/yahoo-client-go/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request
