package main

import (
	"citest/server/core"
	"fmt"
	"net/http"

	"os"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Printf("%+v\n", "run server")
	r := mux.NewRouter()
	r.HandleFunc("/message", core.MessageHandler) // listen message
	r.HandleFunc("/", tes)                        // listen message
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		fmt.Printf("%+v\n", "cannot run server")
	}
}

func tes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "tes GET method")
}
