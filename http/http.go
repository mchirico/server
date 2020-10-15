package http

import (
	"github.com/mchirico/server/http/handles"
	"log"
	"net/http"
)

func SetupHandles() {
	http.HandleFunc("/status", handles.Status)
	http.HandleFunc("/", handles.Status)

}

func Server() {
	SetupHandles()
	log.Println("starting server...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
