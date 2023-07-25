# GoLang & Svelte WebSocket chat
Demo template for a basic working chat application:
- [X] login with a username and delete it when quitting
- [X] create rooms for 1 to 1 or broadcast chat
- [X] send messages to rooms
- [X] add users to rooms

## Server
The server application is built with GoLang, it is an easy server just for demo purpose with an high level web socket implementation.
It uses [this](http://github.com/gobwas/ws) websocket implementation. To start the server simply:
* move to the `server` directory
```sh
$ cd server
```

* run the main program
```sh
$ go run cmf/server0/main.go
```

After that the server should be listening on port 8080 and logs appear in the shell.

### Usage
It is possible to join (connect) and leave (disconnect) from the server.
Once a connection is rooms/chats can be created, messages sent over the rooms/chats and users added to rooms/chats, more details can be found in the next sections.

To test the web socket I've used the following tool: [piesocket websocket tester](https://www.piesocket.com/websocket-tester)

#### Join
To connect and open a websocket conneciton send a request with the websocket tester to the URL:
```sh
ws://localshot:8080/username/foo
```
This will open a connection having a username `foo`, a username cannot be used in more than a connection. To make it available to a different client simply interrupt the connection.

#### Leave
Interrupt the session from the client and the server will handle the disconnection on his own.

#### Create a new room/chat

#### Send a message to a room/chat
#### Add a user to a room/chat
