package chatserver

import (
	"container/list"
	"log"
	"strconv"
	"time"
	"webchat/app/model"
)

type ActiveRoom struct {
	RoomKey   string
	Users     *list.List
	SaveLogs  bool
	Status    bool
	Broadcast chan *Event
}

func NewActiveRoom(rk string) *ActiveRoom {
	activeRoom := &ActiveRoom{
		RoomKey:   rk,
		Users:     list.New(),
		SaveLogs:  true,
		Status:    true,
		Broadcast: make(chan *Event),
	}
	return activeRoom
}

func (r *ActiveRoom) JoinUser(user *OnlineUser) {

	// only one user in a room
	for e := r.Users.Front(); e != nil; e = e.Next() {
		u := e.Value.(*OnlineUser)
		if user.Id == u.Id {
			u.Connection.Close()
			break
		}
	}

	r.Users.PushBack(user)
	// send join message
	event := &Event{
		Type:    "join",
		Text:    user.Info.Name + " has join room",
		User:    user.Info,
		Created: time.Now(),
	}

	log.Println("the room len is:", r.Users.Len())

	r.Broadcast <- event
}

func (r *ActiveRoom) RemoveUser(u *OnlineUser) {
	// remove user form rooms's users list
	for e := r.Users.Front(); e != nil; e = e.Next() {
		user := e.Value.(*OnlineUser)
		if user.Id == u.Id && user.Connection == u.Connection {
			r.Users.Remove(e)
			break
		}
	}
}

func (r *ActiveRoom) UserList() []*UserInfo {
	var userList []*UserInfo

	for u := r.Users.Front(); u != nil; u = u.Next() {
		user := u.Value.(*OnlineUser)
		userList = append(userList, user.Info)
	}
	return userList
}

func (r *ActiveRoom) Run() {
	for {
		select {
		case bc := <-r.Broadcast:
			for u := r.Users.Front(); u != nil; u = u.Next() {
				user := u.Value.(*OnlineUser)
				user.Send <- bc
			}
		}
	}
}

// add user to room recent user list redis
// add user id to set, use Hash save user info
func (r *ActiveRoom) AddUserToRecent(user *model.User) {
	//add userinfo to Hash
	userKey := "user:" + strconv.Itoa(user.Id)

	userInfo := map[string]string{
		"id":     strconv.Itoa(user.Id),
		"name":   user.Name,
		"avatar": user.AvatarUrl(),
	}

	redisClient.Hmset(userKey, userInfo)
	// add user id to room recent user list
	roomKey := "room:" + r.RoomKey + ":users"
	redisClient.Sadd(roomKey, []byte(strconv.Itoa(user.Id)))
}

// get *OnlineUser by user name
func (r *ActiveRoom) GetUserByName(name string) (onlineUser *OnlineUser) {

	for e := r.Users.Front(); e != nil; e = e.Next() {
		user := e.Value.(*OnlineUser)
		if user.Info.Name == name {
			onlineUser = user
			break
		}
	}
	log.Println(onlineUser == nil)
	return onlineUser
}
