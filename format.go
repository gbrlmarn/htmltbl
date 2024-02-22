package main

import (
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

func jsonFormat(tbls []Table) {
    fmt.Println("Needs to be implemented")
}
