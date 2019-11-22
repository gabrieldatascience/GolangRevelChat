package controllers

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/robfig/revel"
	"webchat/app/chatserver"
)

type Websocket struct {
	*revel.Controller
}

func (c Websocket) Chat(roomkey string, ws *websocket.Conn) revel.Result {
	// chech if user has login
	//if !isLogin(c.Controller) {
	//}

	activeRoom := ChatServer.GetActiveRoom(roomkey)
	// crate a user and add usr to room
	user := CurrentUser(c.Controller)
	onlineUser := chatserver.NewOnlineUser(user, ws, activeRoom)
	defer onlineUser.Close()

	// check if user has join room

	activeRoom.JoinUser(onlineUser)
	defer activeRoom.RemoveUser(onlineUser)

	ChatServer.JoinUser(onlineUser)
	defer ChatServer.RemoveUser(onlineUser)

	go onlineUser.PushToClient()

	onlineUser.PullFromClient()

	fmt.Println("websocket.go -the room count is:", ChatServer.ActiveRooms.Len())

	return nil
}
