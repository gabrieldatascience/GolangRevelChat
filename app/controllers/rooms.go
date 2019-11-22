package controllers

import (
	"github.com/robfig/revel"
	"log"
	//"time"
	"webchat/app/chatserver"
	"webchat/app/form"
	"webchat/app/model"
)

type Rooms struct {
	*Application
}

type RoomApi struct {
	*revel.Controller
}

func (c Rooms) Index(p int) revel.Result {
	if p == 0 {
		p = 1
	}

	rooms := model.FindOnePage(p)

	// generate roomlist with recent users
	var roomLists []*model.RoomList
	for _, room := range rooms {
		recentUsers := room.GetRecentUsers() // get []*RecentUser

		rl := &model.RoomList{
			Id:          room.Id,
			RoomKey:     room.RoomKey,
			Title:       room.Title,
			Private:     room.Private,
			Description: room.Description,
			RecentUsers: recentUsers,
		}

		//log.Println("--- the room in rl is:", rl.RoomKey)

		roomLists = append(roomLists, rl)
	}

	allPage := (model.RoomCount() + model.PageSize - 1) / model.PageSize

	return c.Render(p, allPage, roomLists)
}

func (c Rooms) New() revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Sessions.New)
	}

	return c.Render()
}

func (c Rooms) Create(rf *form.RoomForm) revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Sessions.New)
	}

	rf.UserId = CurrentUser(c.Controller).Id

	rf.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Rooms.New)
	}

	room := model.NewRoom(rf)

	if _, err := room.Save(); err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(Rooms.New)
	}

	// run activeroom
	activeroom := chatserver.NewActiveRoom(room.RoomKey)
	go activeroom.Run()
	ChatServer.ActiveRooms.PushBack(activeroom)

	return c.Redirect("/r/%s", room.RoomKey)
}

func (c Rooms) Show(roomkey string) revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Sessions.New)
	}

	currentUser := CurrentUser(c.Controller)

	room := model.FindRoomByRoomKey(roomkey)
	activeRoom := ChatServer.GetActiveRoom(roomkey)
	activeRoom.AddUserToRecent(currentUser)
	log.Println(currentUser.Email)

	// user list
	users := activeRoom.UserList()

	// room list
	rooms := model.FindRoomByUserId(currentUser.Id)
	// user avatar
	userAvatar := currentUser.AvatarUrl()
	// gem latest message
	latestMessages := room.LatestMessage()
	log.Println("latest message len is:", len(latestMessages))

	return c.Render(room, users, userAvatar, rooms, latestMessages)
}

func (c Rooms) Edit(roomkey string) revel.Result {

	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	room := model.FindRoomByRoomKey(roomkey)

	return c.Render(room)
}

func (c Rooms) Update(roomkey string, updateroom *form.UpdateRoom) revel.Result {

	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	room := model.FindRoomByRoomKey(roomkey)

	if err := room.Update(updateroom); err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect("/r/%s/edit", room.RoomKey)
	}

	c.Flash.Success("update success")
	return c.Redirect("/r/%s/edit", room.RoomKey)
}

func (c Rooms) Logs(roomkey string) revel.Result {
	// get all redis from
	logs := model.AllMessageFromRedis(roomkey)

	return c.Render(logs)
}

type UserList struct {
	Users []*chatserver.UserInfo
}

func (c RoomApi) Users(roomkey string) revel.Result {

	// get a activeRoom and get room's user list
	activeroom := ChatServer.GetActiveRoom(roomkey)
	users := activeroom.UserList()

	userList := &UserList{
		Users: users,
	}

	return c.RenderJson(userList)
}
