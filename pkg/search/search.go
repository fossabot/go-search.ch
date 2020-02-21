package search

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	url = "https://tel.search.ch/api/"
)

func (c *client) Search(query ...string) (*Feed, error) {
	client := resty.New()

	was := strings.Join(query, "+")

	res := &Feed{}

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"was": was,
			"key": c.key,
		}).
		SetResult(res).
		SetError(res).
		Get(url)

	if err == nil && resp.StatusCode() != 200 {
		err = fmt.Errorf(resp.Status())
	}

	return res, err
}
