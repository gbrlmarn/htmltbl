package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// Format tables in tabular format
func tableFormat(tbls []Table) {
	for _, t := range tbls {
		table := tablewriter.NewWriter(os.Stdout)
		// Set headers
		table.SetHeader(t.headings)
		for i := 0; i < len(t.rows); i++ {
			table.Append(t.rows[i])
		}
		table.Render()
	}
}

// Map represetation of a HTML table
type mTable map[string][]string

// Create an array of maps representing
// table members
func mkMapTbls(tbls []Table) []mTable {
	var mtbls []mTable
	for _, t := range tbls {
		mt := make(mTable)
		for i := 0; i < len(t.rows); i++ {
			mt[t.headings[i]] = t.rows[i]
		}
		mtbls = append(mtbls, mt)
	}
	return mtbls
}

// Format tables in json format
func jsonFormat(tbls []Table) {
	mtbls := mkMapTbls(tbls)
	jsonTbls, err := json.Marshal(mtbls)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(jsonTbls))
}

// Format tables in json format
func jsonFormatIndent(tbls []Table) {
	mtbls := mkMapTbls(tbls)
	jsonTbls, err := json.MarshalIndent(mtbls, "", "\t")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(jsonTbls))
}
