// Package main implements the tool.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {

	listenAddr := os.Getenv("ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}

	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) { handler(w, r, "post") })
	r.Delete("/", func(w http.ResponseWriter, r *http.Request) { handler(w, r, "delete") })

	log.Printf("listening on %s", listenAddr)

	errListen := http.ListenAndServe(listenAddr, r)

	log.Printf("error: %v", errListen)
}

func handler(w http.ResponseWriter, r *http.Request, method string) {
	body, errBody := io.ReadAll(r.Body)
	if errBody != nil {
		msg := fmt.Sprintf("%s: read body error: %s\n", method, errBody)
		log.Print(msg)
		http.Error(w, msg, 500)
		return
	}
	fmt.Fprintf(w, "%s body:\n", method)
	w.Write(body)
	fmt.Fprintln(w)
}
