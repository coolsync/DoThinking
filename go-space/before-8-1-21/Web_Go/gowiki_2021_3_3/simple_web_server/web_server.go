package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi there, I see %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
