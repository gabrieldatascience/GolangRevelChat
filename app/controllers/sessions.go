package controllers

import (
	"github.com/robfig/revel"
	"webchat/app/form"
	"webchat/app/model"
	//"fmt"
)

type Sessions struct {
	*Application
}

func (c Sessions) New() revel.Result {
	//fmt.Println(c.Session["user_name"])
	return c.Render()
}

func (c Sessions) Create(loginform *form.UserLogin) revel.Result {
	loginform.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Sessions.New)
	}

	if !model.Authenticate(loginform.Name, loginform.Password) {
		c.Flash.Error("username or password error")
		return c.Redirect(Sessions.New)
	}

	//create session
	c.Session["user_name"] = loginform.Name
	c.Flash.Success("Login success")
	return c.Redirect(Rooms.Index)
}

func (c Sessions) Destroy() revel.Result {
	user := CurrentUser(c.Controller)

	for k := range c.Session {
		delete(c.Session, k)
	}

	if onlineUser := ChatServer.GetUserById(user.Id); onlineUser != nil {
		// close conn
		onlineUser.Connection.Close()
	}

	return c.Redirect(Application.Index)
}
