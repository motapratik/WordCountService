package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/motapratik/WordCountService/api"
)

func main() {
	//create a new router
	router := mux.NewRouter()
	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/TopTenWordCount", api.TopTenWordCount).Methods("POST")
	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
