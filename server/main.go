package main

import (
	"citest/server/core"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Printf("%+v\n", "run server")
	r := mux.NewRouter()
	r.HandleFunc("/message", core.MessageHandler) // listen message
	r.HandleFunc("/", tes)                        // listen message
	err := http.ListenAndServe(":5000", r)
	if err != nil {
		fmt.Printf("%+v\n", "cannot run server")
	}
}

func tes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "tes GET method")
}
