package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Format tables in tabular format
func tableFormat(tbls []Table) {
	for _, t := range tbls {
		table := tablewriter.NewWriter(os.Stdout)
		// Set headers
		table.SetHeader(t.ths)
		for i := 0; i < len(t.tds); i += len(t.ths) {
			table.Append(t.tds[i : i+len(t.ths)])
		}

		table.Render()
	}
}

