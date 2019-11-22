package chatserver

import (
	"container/list"
	"time"
	//"errors"
	"log"
	"webchat/app/model"
)

type Server struct {
	Name        string
	ActiveRooms *list.List
	ActiveUsers *list.List
}

type Event struct {
	Type    string
	Text    string
	Created time.Time
	User    *UserInfo
}

func NewServer() *Server {
	Fx := &Server{
		Name:        "webchat",
		ActiveRooms: list.New(),
		ActiveUsers: list.New(),
	}
	return Fx
}

// find avtive room return a activeroom instance
func (s *Server) GetActiveRoom(roomkey string) *ActiveRoom {
	var activeroom *ActiveRoom
	for room := s.ActiveRooms.Front(); room != nil; room = room.Next() {
		r := room.Value.(*ActiveRoom)
		if r.RoomKey == roomkey {
			activeroom = r
		}
	}

	if activeroom == nil {
		activeroom = NewActiveRoom(roomkey)
		go activeroom.Run()
		s.ActiveRooms.PushBack(activeroom)
	}

	return activeroom
}

// Get all run rooms
func (s *Server) AllRunRooms() []*ActiveRoom {
	var rooms []*ActiveRoom
	for room := s.ActiveRooms.Front(); room != nil; room = room.Next() {
		r := room.Value.(*ActiveRoom)
		rooms = append(rooms, r)
	}
	return rooms
}

// init all room
func (s *Server) RunRooms() {
	rooms := model.AllRoom()

	for _, room := range rooms {
		activeroom := NewActiveRoom(room.RoomKey)
		// run room in a goroutine and push room to list
		go activeroom.Run()
		s.ActiveRooms.PushBack(activeroom)
	}
}

func (s *Server) JoinUser(u *OnlineUser) {
	s.ActiveUsers.PushBack(u)
	log.Println("the server user list len is:", s.ActiveUsers.Len())
}

func (s *Server) RemoveUser(u *OnlineUser) {
	// remove user from server users list
	for e := s.ActiveUsers.Front(); e != nil; e = e.Next() {
		user := e.Value.(*OnlineUser)
		if user.Id == u.Id && user.Connection == u.Connection {
			s.ActiveUsers.Remove(e)
			break
		}
	}

	log.Println("the server user list len is:", s.ActiveUsers.Len())
}

// get *OnlineUser by id
func (s *Server) GetUserById(id int) *OnlineUser {
	var onlineUser *OnlineUser

	for e := s.ActiveUsers.Front(); e != nil; e = e.Next() {
		user := e.Value.(*OnlineUser)
		if user.Id == id {
			onlineUser = user
			break
		}
	}
	return onlineUser
}
