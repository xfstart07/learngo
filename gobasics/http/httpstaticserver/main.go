package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Print("Listening on :3333...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
