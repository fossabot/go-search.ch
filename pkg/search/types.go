package search

import "encoding/xml"

// Client search client interface
type Client interface {
	Search(query ...string) (*Feed, error)
}

// client implements Client
var _ Client = &client{}

type client struct {
	key string
}

// Feed result feed
type Feed struct {
	XMLName      xml.Name  `xml:"feed"`
	Lang         string    `xml:"lang,attr"`
	Xmlns        string    `xml:"xmlns,attr"`
	OpenSearch   string    `xml:"openSearch,attr"`
	Tel          string    `xml:"tel,attr"`
	ID           string    `xml:"id"`
	Title        Type      `xml:"title"`
	Generator    Generator `xml:"generator"`
	Updated      string    `xml:"updated"`
	Link         []Link    `xml:"link"`
	ErrorCode    string    `xml:"errorCode,omitempty"`
	ErrorReason  string    `xml:"errorReason,omitempty"`
	ErrorMessage string    `xml:"errorMessage,omitempty"`
	TotalResults string    `xml:"totalResults"`
	StartIndex   string    `xml:"startIndex"`
	ItemsPerPage string    `xml:"itemsPerPage"`
	Query        Query     `xml:"Query"`
	Entry        []Entry   `xml:"entry"`
}

// Link link element
type Link struct {
	Href  string `xml:"href,attr"`
	Title string `xml:"title,attr"`
	Rel   string `xml:"rel,attr"`
	Type  string `xml:"type,attr"`
}

// Type type element
type Type struct {
	Type string `xml:"type,attr"`
}

// Author author element
type Author struct {
	Name string `xml:"name"`
}

// Query query element
type Query struct {
	Role        string `xml:"role,attr"`
	SearchTerms string `xml:"searchTerms,attr"`
	StartPage   string `xml:"startPage,attr"`
}

// Entry result entry
type Entry struct {
	ID         string   `xml:"id"`
	Updated    string   `xml:"updated"`
	Published  string   `xml:"published"`
	Title      Type     `xml:"title"`
	Content    Type     `xml:"content"`
	Autor      Author   `xml:"autor"`
	Link       []Link   `xml:"link"`
	Pos        string   `xml:"pos"`
	Type       string   `xml:"type"`
	Name       string   `xml:"name"`
	Firstname  string   `xml:"firstname"`
	Occupation string   `xml:"occupation"`
	Street     string   `xml:"street"`
	Streetno   string   `xml:"streetno"`
	Zip        string   `xml:"zip"`
	City       string   `xml:"city"`
	Canton     string   `xml:"canton"`
	Phone      string   `xml:"phone"`
	Category   []string `xml:"category"`
	Extra      []Type   `xml:"extra"`
}

// Generator generator
type Generator struct {
	Version string `xml:"version,attr"`
	URI     string `xml:"uri,attr"`
}
