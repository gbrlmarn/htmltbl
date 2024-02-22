package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
)

type Table struct {
	th   []string
	td   []string
	cols int
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please specify the link for table extraction")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Only one link is supported")
		return
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	tbls, err := Extract(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
    Render(tbls)
}

func Extract(body io.ReadCloser) ([]Table, error) {
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

func Render(tbls []Table) {
	table := tablewriter.NewWriter(os.Stdout)
	// Set headers
	for _, t := range tbls {
		table.SetHeader(t.th)
        table.Append(t.td)
        table.Render()
	}
}
