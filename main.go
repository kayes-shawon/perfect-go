package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"errors"
	"github.com/auth0/go-jwt-middleware"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
)


type Product struct {
	Id int
	Name string
	Slug string
	Description string
}

var products = []Product{
	{
		Id:          1,
		Name:        "World of Authcraft",
		Slug:        "world-of-authcraft",
		Description: "Battle bugs and protect yourself from invaders while you explore a scary world with no security",
	},
	{
		Id:          2,
		Name:        "Ocean Explorer",
		Slug:        "ocean-explorar",
		Description: "Explore the depths of the sea in this one of a kind underwater experience",
	},
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API up and running"))
})

func main() {

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options {

	})

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views")))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/products", jwtMiddleware.Handler(ProductHandler)).Methods("GET")
	r.Handle("/products/{slug}/feedback", jwtMiddleware.Handler(AddFeedbackHandler)).Methods("POST")


	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":8080", r)
}

var ProductHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([] byte(payload))
})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

