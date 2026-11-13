package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You`ve requested the book: %s on page %s\n", title, page)
}

func main() {
	r := mux.NewRouter()

	bookrouter := r.PathPrefix("/books").Subrouter()

	bookrouter.HandleFunc("/{title}/page/{page}", ReadBook).Methods("GET")

	http.ListenAndServe(":5555", r)
}
