package core

import (
	"log"
	"fmt"
	"net/http"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: ", r.URL.Path)

	fmt.Fprintf(w, "Hello World 4!")
}
