package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting to listen on port 8080")
	http.ListenAndServe("0.0.0.0:8080", http.FileServer(http.Dir("./docs")))
}
