package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	igaccountsearch "github.com/codenoid/instagram-account-search"
)

var bind string

func main() {
	flag.StringVar(&bind, "bind", ":8080", "bind 127.0.0.1:8080")
	flag.Parse()

	http.HandleFunc("/search", searchIGAccount)
	fmt.Println("starting server at", bind)
	log.Fatal(http.ListenAndServe(bind, nil))
}

func searchIGAccount(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	w.Header().Add("Content-Type", "application/json")
	searchResult, err := igaccountsearch.UserSearch(q)
	if err != nil {
		fmt.Fprintf(w, `{"error": true, "message": %v}`, err.Error())
		return
	}
	b, err := json.Marshal(searchResult)
	if err != nil {
		fmt.Fprintf(w, `{"error": true, "message": %v}`, err.Error())
		return
	}
	w.Write(b)
}
