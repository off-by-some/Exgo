package socket

import (
  http "net/http"
  "github.com/gorilla/websocket"
)



var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// TODO: Require auth for connecting
// The session cookie would be sent
func Connect(res http.ResponseWriter, req *http.Request) {
  // TODO: Care about this error
  conn, _ := upgrader.Upgrade(w, r, nil)
}


// Every node that owns a room is connected
// to the server over websocket.
// New nodes connecting GET to the server with
// the host's ID, the server then holds that
// request open while it sends a message to the host
// node requesting SDP info, which it then sends
// back to the new node and closes the request.
