package main

import (
	"fmt"
	mux "github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/cache", CacheHandler).Methods("GET")

	http.Handle("/", router)

	fmt.Println("Handling it.")
	http.ListenAndServe(":8000", router)

}

func CacheHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cache & release"))
}
