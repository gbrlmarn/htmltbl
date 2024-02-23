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
		table.SetHeader(t.th)
		for i := 0; i < len(t.td); i += t.cols {
			table.Append(t.td[i : i+t.cols])
		}
		table.Render()
	}
}

// Map representation of a table row
type mrow map[string]string

// Map representation of a table
type mtable []mrow 

// Create an array of maps representing
// table members
func mkMapTbls(tbls []Table) []mtable {
    var mtbls []mtable
	for _, t := range tbls {
        var mt mtable 
		for i := 0; i < len(t.td); i += t.cols {
            mr := make(mrow, t.cols)
			for j := 0; j < t.cols; j++ {
				mr[t.th[j]] = t.td[i+j]
			}
            mt = append(mt, mr)
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

