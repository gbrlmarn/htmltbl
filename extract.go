package main

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

// Table struct containing 'th' tags,
// 'td' tags and number of cols
type Table struct {
	th   []string
	td   []string
	cols int
}

// Extract all tables of tag '<table>' from 
// a html body reader closer
func extract(body io.ReadCloser) ([]Table, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}
	var tables []Table
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		var t Table
		// find table header
		s.Find("th").Each(func(i int, s *goquery.Selection) {
			t.th = append(t.th, s.Text())
			t.cols += 1
		})
		// find table data
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			t.td = append(t.td, s.Text())
		})
		tables = append(tables, t)
	})
	return tables, nil
}

