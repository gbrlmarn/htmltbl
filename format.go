package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

// Format tables in tabular format
func tableFormat(tbls []Table) {
	for _, t := range tbls {
		table := tablewriter.NewWriter(os.Stdout)
		// Set headers
		table.SetHeader(t.Headers)
		for i := 0; i < len(t.Data); i += len(t.Headers) {
			table.Append(t.Data[i : i+len(t.Headers)])
		}

		table.Render()
	}
}

// Format tables in json format
func jsonFormat(tbls []Table) {
	jsonTbls, err := json.Marshal(tbls)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(jsonTbls))
}

// Format tables in json format
func jsonIndentFormat(tbls []Table) {
	jsonTbls, err := json.MarshalIndent(tbls, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(jsonTbls))
}
