package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("r.Header.Get(\"User-Agent\"): %v\n", r.Header.Get("User-Agent"))
		fmt.Fprintln(w, "Your using: ", r.Header.Get("User-Agent"))
	})

	http.ListenAndServe(":80", nil)
}
