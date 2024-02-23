package main

import (
	"net/http"
)

// Fetch html body from url and 
// returns an array of tables
func fetch(url string) ([]Table, error) {
	resp, err := http.Get(url)
	if err != nil {
        return nil, err
	}
	defer resp.Body.Close()

	tbls, err := extract(resp.Body)
	if err != nil {
        return nil, err
	}
	return tbls, nil 
}
