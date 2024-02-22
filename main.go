package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
    format = flag.String("format", "table", "format: table, json")
)

func main() {
    flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Please specify the link for table extraction")
		return
	}
	if len(flag.Args()) > 2 {
		fmt.Println("Only one link is supported")
		return
	}
    url := flag.Args()[0]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	tbls, err := extract(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

    switch *format {
    case "table":
        tableFormat(tbls)
    case "json":
        jsonFormat(tbls)
    default:
        panic(fmt.Sprintf("unsupported format %s\n", *format)) 
    }
}

