package main

import (
	"fmt"
	"net/http"
)

func main() {
	const numStr string = "2"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! "+numStr)
	})

	fmt.Println("Server starting on port 8080...")
	fmt.Println(numStr)
	http.ListenAndServe(":8080", nil)
}
