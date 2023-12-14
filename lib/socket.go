package lib

import (
	socketio "github.com/googollee/go-socket.io"
)

func Socket() *socketio.Server {
	server := socketio.NewServer(nil)
	return server
}
