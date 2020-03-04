package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"html"
)

func main() {

	handle := func(w http.ResponseWriter, r *http.Request) {
		log.Print("Received Request")

		// simulate long backend action, e.g. db query
		time.Sleep(5 * time.Second)

		fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
		log.Print("Sent Response")
	}

	http.HandleFunc("/", handle)
	http.HandleFunc("/1", handle)
	http.HandleFunc("/2", handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
