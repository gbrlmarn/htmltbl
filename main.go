package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
    format = flag.String("format", "table", "format: table")
)

func main() {
    flag.Parse()
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

	tbls, err := extract(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

    switch *format {
    case "table":
        render(tbls)
    default:
        panic(fmt.Sprintf("unsupported format %s", *format)) 
    }
}

