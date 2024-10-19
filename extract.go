package main

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Table struct extracted from HTML
type Table struct {
	ths  []string
	tds  []string
}

// Extract all tables of tag '<table>' from
// a html body reader closer
func extract(tbls []Table, n *html.Node) []Table {
	switch n.Data {
	case "table":
		tbls = append(tbls, Table{})
	case "th":
		var sb strings.Builder
		innerText(n, &sb)
		tbls[len(tbls)-1].ths = append(tbls[len(tbls)-1].ths, sb.String())
	case "td":
		var sb strings.Builder
		innerText(n, &sb)
		tbls[len(tbls)-1].tds = append(tbls[len(tbls)-1].tds, sb.String())
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tbls = extract(tbls, c)
	}
	return tbls
}

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
	tbls = extract(tbls, doc)
	if err != nil {
		return nil, err
	}
	return tbls, nil
}
