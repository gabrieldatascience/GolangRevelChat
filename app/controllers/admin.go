package controllers

import (
	"github.com/robfig/revel"
	"log"
	"webchat/app/form"
	"webchat/app/model"
)

type Admin struct {
	*Application
}

func (c Admin) Index() revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

	roomCount := ChatServer.ActiveRooms.Len()
	onlineUserCount := ChatServer.ActiveUsers.Len()
	latestUsers := model.LatestUsers(5)
	userCount := model.UserCount()
	ServerName := ChatServer.Name

	return c.Render(roomCount, onlineUserCount, latestUsers, userCount, ServerName)
}

func (c Admin) Users() revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

	users := model.AllUsers()
	return c.Render(users)
}

func (c Admin) Rooms() revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

	rooms := ChatServer.AllRunRooms()
	return c.Render(rooms)

}

func (c Admin) ChangeLogStatus(roomKey string) revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

	room := ChatServer.GetActiveRoom(roomKey)
	room.SaveLogs = !room.SaveLogs
	return c.Redirect(Admin.Rooms)
}

func (c Admin) ChangeRunStatus(roomKey string) revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

	room := ChatServer.GetActiveRoom(roomKey)
	room.Status = !room.Status
	return c.Redirect(Admin.Rooms)
}

func (c Admin) SiteSettings() revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

	return c.Render()
}

func (c Admin) SaveServerSettings(sf *form.ServerSettings) revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

	sf.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Admin.SiteSettings)
	}

	log.Println("form: ", sf.Name)

	ChatServer.Name = sf.Name

	return c.Redirect(Admin.SiteSettings)
}

func (c Admin) OnlineUsers(roomKey string) revel.Result {
	room := ChatServer.GetActiveRoom(roomKey)
	users := room.UserList()

	return c.Render(users)
}

func (c Admin) checkAdmin() bool {
	user := model.FindUserByName(c.Session["user_name"])
	if user.Email == "ldshuang@gmail.com" {
		return true
	} else {
		return false
	}
}
