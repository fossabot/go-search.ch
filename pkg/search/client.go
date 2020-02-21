package search

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
	"golang.org/x/text/encoding/charmap"
)

const (
	// DefaultURL the default api service url
	DefaultURL = "https://tel.search.ch/api/"
	// TestURL the test api service url
	TestURL = "https://tel.search.ch/examples/api-response.xml"
)

// Client search client interface
type Client interface {
	Search(query ...string) (*Feed, error)
}

// New create a new client with the given key
func New(key string) Client {
	return NewFor(Config{
		Key: key,
	})
}

// NewFor create a new client for the given config
func NewFor(config Config) Client {
	url := config.URL
	if url == "" {
		url = DefaultURL
	}

	r := resty.New().
		SetHeader("Accept", "text/xml")

	if config.InsecureSkipVerify {
		r.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	return &client{
		key: config.Key,
		url: url,
		r:   r,
	}
}

// Config client config
type Config struct {
	// Key the search.ch API key
	Key string
	// URL the service URL
	URL string
	// InsecureSkipVerify ignore tls
	InsecureSkipVerify bool
}

// client implements Client
var _ Client = &client{}

type client struct {
	key string
	url string
	r   *resty.Client
}

func (c *client) Search(query ...string) (*Feed, error) {

	was := strings.Join(query, "+")

	res := &Feed{}

	resp, err := c.r.R().
		SetQueryParams(map[string]string{
			"was": was,
			"key": c.key,
		}).
		Get(c.url)

	if err == nil && resp.StatusCode() != 200 {
		err = fmt.Errorf(resp.Status())
	}

	parseXML(resp.Body(), res)

	return res, err
}

func parseXML(xmlDoc []byte, target interface{}) {
	reader := bytes.NewReader(xmlDoc)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = makeCharsetReader
	if err := decoder.Decode(target); err != nil {
		log.Fatalf("unable to parse XML '%s':\n%s", err, xmlDoc)
	}
}

func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "ISO-8859-1" {
		// Windows-1252 is a superset of ISO-8859-1, so should do here
		return charmap.Windows1252.NewDecoder().Reader(input), nil
	}
	return nil, fmt.Errorf("Unknown charset: %s", charset)
}
