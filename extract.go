package main

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Table struct extracted from HTML
type Table struct {
	Headers []string
	Data    []string
	Ncols   int
	Nrows   int
}

// Extract all tables of tag '<table>' from
// a html body reader closer
func extract(tbls []Table, n *html.Node) ([]Table, error) {
	switch n.Data {
	case "table":
		tbls = append(tbls, Table{})
	case "tr":
		tbls[len(tbls)-1].Nrows += 1
	case "th":
		var sb strings.Builder
		innerText(n, &sb)
		tbls[len(tbls)-1].Headers = append(tbls[len(tbls)-1].Headers, sb.String())
		tbls[len(tbls)-1].Ncols += 1
	case "td":
		var sb strings.Builder
		innerText(n, &sb)
		tbls[len(tbls)-1].Data = append(tbls[len(tbls)-1].Data, sb.String())
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		var err error
		tbls, err = extract(tbls, c)
		if err != nil {
			return nil, err
		}
	}
	return tbls, nil
}

// Extract inner text from html nodes
func innerText(n *html.Node, sb *strings.Builder) {
	if n.Type == html.TextNode {
		sb.WriteString(strings.TrimSpace(n.Data))
		return
	}
	if n.FirstChild == nil {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		innerText(c, sb)
	}
}

// Fetch html body from url and
// returns an array of tables
func fetch(url string) ([]Table, error) {
	var tbls []Table
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	tbls, err = extract(tbls, doc)
	if err != nil {
		return nil, err
	}
	return tbls, nil
}
