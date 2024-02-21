package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
    fmt.Println(string(body))
}
