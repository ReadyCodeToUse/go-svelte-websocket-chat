package room

import (
	"errors"
	"sync"

	"github.com/cedrata/go-svelte-websocket-chat/server/pkg/user"
	"github.com/gobwas/ws"
)

type Room struct {
	users *[]user.User
	name  string
	*sync.RWMutex
}

type RoomsList struct {
	list map[string]*Room
	*sync.RWMutex
}

// Returns a new room with the specified name
func NewRoom(name string) *Room {
	return &Room{&[]user.User{}, name, &sync.RWMutex{}}
}

// Returns the value of the rooms name
func (r Room) GetName() string {
	r.RLock()
	defer r.RUnlock()
	return r.name
}

// Add a new user to the rooms instance if it does not exists,
// if the user already exists returns false.
func (r *Room) AddUser(u *user.User) bool {
	exist := false
	name := u.GetName()

	r.RLock()
	for _, item := range *r.users {
		if item.GetName() == name {
			exist = true
			break
		}
	}
	r.RUnlock()

	if exist {
		return false
	}

	r.Lock()
	*r.users = append(*r.users, *u)
	defer r.Unlock()

	return true
}

// Sends a given message to the specified room
func (r *Room) SendMessage(msg, sender string, opCode ws.OpCode) error {
	bmsg := []byte(msg)
	var errs error

	r.RLock()
	for _, item := range *r.users {
		if item.GetName() == sender {
			continue
		}

		err := item.Receive(bmsg, opCode)
		errs = errors.Join(errs, err)
	}
	r.RUnlock()
	return errs
}

// Returns a new empty rooms list
func NewRoomsList() *RoomsList {
	return &RoomsList{make(map[string]*Room), &sync.RWMutex{}}
}

// Adds a room to the list, if added correctly returns true
func (rl *RoomsList) AddRoom(r *Room) bool {
	rl.Lock()
	defer rl.Unlock()

	_, ok := rl.list[r.GetName()]
	if ok {
		return false
	}

	rl.list[r.GetName()] = r
	return true
}

// Deletes a room from the list, if the room does not exists nothing happens.
func (rl *RoomsList) DeleteRoom(r *Room) {
	rl.Lock()
	defer rl.Unlock()

	delete(rl.list, r.GetName())
}

// Get a room from the list, if the room does not exists returns false as second return value
func (rl *RoomsList) GetRoom(name string) (*Room, bool) {
	rl.RLock()
	defer rl.RUnlock()

	res, ok := rl.list[name]
	return res, ok
}
