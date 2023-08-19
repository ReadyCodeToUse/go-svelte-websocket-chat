# GoLang & Svelte WebSocket chat
Demo template for a basic working chat application:
- [X] login with a username and delete it when quitting
- [X] create rooms for 1 to 1 or broadcast chat
- [X] send messages to rooms
- [X] add users to rooms

## Server
The server application is built with GoLang, it is an easy server just for demo purpose with an high level web socket implementation.
It uses [this](http://github.com/gobwas/ws) websocket implementation. To start the server simply:
move to the `server` directory
```sh
$ cd server
```

run the main program
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

The username foreach connection can only be a lowercase alphanumeric string.

#### Leave
Interrupt the session from the client and the server will handle the disconnection on his own.

#### Create a new room/chat
When the connection is estabilshed a message with the following fields can be sent:
```json
{
    "action": "create-room",
    "room": "%room-name%", // required
    "content": "", // ignored
    "users": ["%some-username%", "%some-more-username%"], // optional
}
```

`action` has to be equal to `create-room` in order to enable the action of creating a new room. A new room must have a name and must be provided in the `room` field. `content` is ignored with this action and last a list of users can be provided to create a not empty room.

#### Send a message to a room/chat
When the connection is estabilshed a message with the following fields can be sent:
```json
{
    "action": "send-msg",
    "room": "%room-name%", // required
    "content": "%some-message-content%", // requred
    "users": [], // ignored
}
```

`action` has to be equal to `seng-msg` in order to enable the action of sending a message. A room must be provided in the `room` field to know what room to send the message too. `content` must have a value other than empty string and last the list of users will be ignored by default.
#### Add a user to a room/chat
See [create a new room/chat](#create-a-new-roomchat) but instead of `create-room` in the `action` field write `add-user`

## Client
DISCLAIMER: this is not a complete working client, it was just an exercise for me to learn and improve my frontend skills and knowledge and this is a report on what I've done and what I did correctly or wrong in my opinion. This project can be connected to the backend with the use of web sockets (not done for now becuase I'd like to focus on something else now) and also contains some commented code that shows an alert box before closing the tab in order to control an eventual websocket disconnection.

The informations on how to start the project are available [here](/client/README.md).

### Good
Imrpoved knowledge about componenets in general, componenet events and slots.
Using a state machine (not really the best use case) to open a modal box, to learn how to write it with Svelte stores (if something more comples XState can be used)
Tailwind styling in general.
Writing some JS code.

### Bad/Wrong
The main thing I would do differently is how the managing of the previews, an array named `previews` is created and populated when the components are created. That is not a correct behaviour IMHO, a better data struture contianing all the informations about the rooms, messages and previews could be used and that would make easier to manage the creation of the previews. In order to select the newly created preview a `setTimeout` has been added in the `+page.svelte`. Doing so previews has time to trigger the svelte each loop and after the timeout the new preview is selected.