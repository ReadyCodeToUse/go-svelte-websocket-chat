package main

import (
	"log"
	"net/http"

	"github.com/cedrata/go-svelte-websocket-chat/server/pkg/connection"
)

func main() {
	log.Println("starting server listening on poirt 8080")

	mux := http.NewServeMux()
	mux.Handle("/username/", connection.NewConnectionHandler())
	http.ListenAndServe(":8080", mux)

	log.Panicln("Server closed")
}
