package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/cedrata/go-svelte-websocket-chat/server/pkg/room"
	"github.com/cedrata/go-svelte-websocket-chat/server/pkg/user"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Action string

const (
	SEND_MSG    Action = "send-msg"    // Send a message over a room
	CREATE_ROOM Action = "create-room" // Create a new room
	ADD_USERS   Action = "add-users"   // Add user to a room
)

// Struct to hold the message coming from the client. The client
// will send data according to the requried action:
//   - action = 'send-msg' must have room and content valorized, users field will be ignored
//   - action = 'create-room' will create a room with 1 user by default
//     (request sender) and users provided in
//     users array, content will be ignored
//   - action = 'add-user' requires the room and users field to be filled, content is ignored
type RequestPayload struct {
	Action  Action   `json:"action"`
	Room    string   `json:"room"`
	Content string   `json:"content"`
	Users   []string `json:"users"`
}

// Struc able to handle a new client connection
type ConnectionHandler struct {
	ul user.UsersList
	rl room.RoomsList
}

// Create a new connection handle with empty users and rooms list
func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{
		ul: *user.NewUsersList(),
		rl: *room.NewRoomsList(),
	}
}

// Parses a message and return the generic request payload
func ParseMessage(msg []byte) *RequestPayload {
	var res RequestPayload
	json.Unmarshal(msg, &res)
	return &res
}

func getEndpointRegex() *regexp.Regexp {
	return regexp.MustCompile(`^\/username\/([a-z0-9]+)$`)
}

func (ch *ConnectionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && getEndpointRegex().MatchString(r.URL.Path):
		username := strings.TrimPrefix(r.URL.Path, "/username/")

		if ch.ul.UsernameExists(username) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("The selected user is already logged in, cannot accept incoming connection"))
			log.Printf("User %s is already connected, terminating connection", username)
			return
		}

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An error occured during the creation of the web socket connection"))
			return
		}

		user := user.NewUser(username, conn)
		ch.ul.AddUser(user)
		go ch.serve(user)

		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}
}

// Serve web socket for provided user.
func (ch *ConnectionHandler) serve(u *user.User) {
	defer func() {
		log.Println("Closing connection")
		if err := u.CloseConnection(); err != nil {
			log.Println(err)
		}
		ch.ul.DeleteUser(u)
	}()

	conn := u.GetConnection()

	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			log.Println("First loop")
			log.Println(err)
			return
		}

		req := RequestPayload{}
		err = json.Unmarshal(msg, &req)
		if err != nil {
			wsutil.WriteServerMessage(conn, op, []byte(err.Error()))
		}

		switch req.Action {
		case SEND_MSG:
			ch.sendMessage(&req, u.GetName(), op)
		case CREATE_ROOM:
			ch.createRoom(&req)
		case ADD_USERS:
			ch.addUsers(&req)
		default:
		}

		err = wsutil.WriteServerMessage(conn, op, msg)
		if err != nil {
			log.Println("Second loop")
			log.Println(err)
			return
		}
	}
}

func (ch *ConnectionHandler) sendMessage(p *RequestPayload, sender string, op ws.OpCode) error {
	r, ok := ch.rl.GetRoom(p.Room)
	if !ok {
		return fmt.Errorf("the requested room does not exists")
	}

	err := r.SendMessage(p.Content, sender, op)
	return err
}

func (ch *ConnectionHandler) createRoom(p *RequestPayload) error {
	r := room.NewRoom(p.Room)

	if ok := ch.rl.AddRoom(r); !ok {
		return errors.New("the choosen room name is already in use, select a different one")
	}

	err := ch.addUsers(p)
	return err
}

func (ch *ConnectionHandler) addUsers(p *RequestPayload) error {
	r, ok := ch.rl.GetRoom(p.Room)
	if !ok {
		return fmt.Errorf("the requested room does not exists")
	}

	var err error = nil

	users := []*user.User{}
	for _, item := range p.Users {
		user, ok := ch.ul.GetUser(item)

		if ok {
			users = append(users, user)
		} else {
			err = errors.Join(err, fmt.Errorf("the user %s does not exists", item))
		}
	}
	if err != nil {
		return err
	}

	for _, item := range users {
		ok := r.AddUser(item)
		if !ok {
			err = errors.Join(err, fmt.Errorf("the user named %s aleady exists in this room, user not added", item.GetName()))
		}
	}
	return err
}
