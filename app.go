package main

import (
	"fetcher/fetcher"
	"os"
)

func main() {
	f := fetcher.NewFetchSite(os.Args[1])
	f.Fetch()
}
