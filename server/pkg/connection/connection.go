package connection

import "net/http"

type ConnectionHandler struct {
}

func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{}
}

func (ch *ConnectionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
