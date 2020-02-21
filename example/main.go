package main

import (
	"fmt"

	"github.com/bakito/go-search.ch/pkg/search"
)

func main() {
	cl := search.NewFor(search.Config{
		Key: "xxx",
		URL: search.TestURL,
	})

	res, err := cl.Search("0111111111")
	if err != nil {
		panic(err)
	}

	for _, e := range res.Entry {
		fmt.Println(e)
	}

}
