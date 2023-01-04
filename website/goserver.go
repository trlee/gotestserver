package main

import (
	"net/http"
	"websitetest/website/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
