package main

import (
	"net/http"

	"github.com/LostLaser/Labeloos/views"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", views.HomeView)

	http.ListenAndServe(":8080", r)
}
