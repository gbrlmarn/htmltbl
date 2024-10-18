package main

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

// Table struct extracted from HTML
type Table struct {
	headings []string
	rows     [][]string
}

// Extract all tables of tag '<table>' from
// a html body reader closer
func extract(body io.ReadCloser) ([]Table, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}
	var tbls []Table
	doc.Find("table").Each(func(i int, tablehtml *goquery.Selection) {
		tbls = append(tbls, extractTbl(tablehtml))
	})
	return tbls, nil
}

func extractTbl(tblhtml *goquery.Selection) Table {
	var tbl Table
	var row []string

	// find table row
	tblhtml.Find("tr").Each(func(i int, rowhtml *goquery.Selection) {
		// find table header
		rowhtml.Find("th").Each(func(i int, tableheading *goquery.Selection) {
			tbl.headings = append(tbl.headings, tableheading.Text())
		})
		// find table data
		rowhtml.Find("td").Each(func(i int, tablecell *goquery.Selection) {
			row = append(row, tablecell.Text())
		})
		tbl.rows = append(tbl.rows, row)
		row = nil
	})
	return tbl
}
