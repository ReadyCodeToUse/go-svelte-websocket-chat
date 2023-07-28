package user

import (
	"fmt"
	"net"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type User struct {
	name string
	conn net.Conn
	*sync.RWMutex
}

type UsersList struct {
	list map[string]*User
	sync.RWMutex
}

// Returns a new user with the specified username and connection
func NewUser(name string, conn net.Conn) *User {
	return &User{name, conn, &sync.RWMutex{}}
}

// Returns the value of the users name
func (u User) GetName() string {
	u.RLock()
	defer u.RUnlock()
	return u.name
}

// Returns the value of the users connection
func (u User) GetConnection() net.Conn {
	u.RLock()
	defer u.RUnlock()
	return u.conn
}

// Close an open connection
func (u *User) CloseConnection() error {
	return u.conn.Close()
}

// Send a message to the user with his channel.
func (u *User) Receive(msg []byte, opCode ws.OpCode) error {
	u.Lock()
	defer u.Unlock()
	err := wsutil.WriteServerMessage(u.conn, opCode, msg)
	if err != nil {
		return err
	}

	return nil
}

// Return a new empty user list
func NewUserList() *UsersList {
	return &UsersList{make(map[string]*User), sync.RWMutex{}}
}

// Returns true if the username exists in the UserList instance
func (ul *UsersList) UsernameExists(username string) bool {
	ul.RLock()
	defer ul.RUnlock()

	_, ok := ul.list[username]

	return ok
}

// Adds a user to the user list, if the user already exists it will be replaced
func (ul *UsersList) AddUser(u *User) {
	ul.Lock()
	defer ul.Unlock()

	ul.list[u.GetName()] = u
}

// Remove a user from the user list, if the user does not exists nothing happends
func (ul *UsersList) DeleteUser(u *User) {
	ul.Lock()
	defer ul.Unlock()

	delete(ul.list, u.GetName())
}

func (ul *UsersList) GetUser(name string) (*User, error) {
	ul.RLock()
	defer ul.RUnlock()

	res, ok := ul.list[name]
	if !ok {
		return nil, fmt.Errorf("the user named %s does not exists", name)
	}

	return res, nil
}
