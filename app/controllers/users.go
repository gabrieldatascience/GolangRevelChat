package controllers

import (
	"github.com/robfig/revel"
	"webchat/app/form"
	"webchat/app/model"
	//"mime/multipart"
	// "log"
)

type Users struct {
	*Application
}

func (c Users) New() revel.Result {
	return c.Render()
}

func (c Users) Create(userform *form.UserForm) revel.Result {
	userform.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Users.New)
	}

	user := model.NewUser(userform)
	err := user.Save()

	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(Users.New)
	}

	c.Flash.Success("signup success, please login")
	return c.Redirect(Sessions.New)
}

func (c Users) MyRooms() revel.Result {
	user := CurrentUser(c.Controller)
	rooms := user.Rooms()

	return c.Render(rooms)
}

func (c Users) EditSettings() revel.Result {
	user := CurrentUser(c.Controller)

	return c.Render(user)
}

func (c Users) SaveSettings(setting *form.Settings) revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	user := CurrentUser(c.Controller)

	if err := user.SaveSettings(setting); err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(Users.EditSettings)
	}

	c.Flash.Success("save success")

	return c.Redirect(Users.EditSettings)
}

// func (c Users) Avatar(file []byte) revel.Result {
// 	if !isLogin(c.Controller) {
// 		c.Flash.Error("Please login first")
// 		return c.Redirect(Application.Index)
// 	}
//     avatar := os.Open()

//     log.Println("file is:", file)
// 	return nil
// }

func (c Users) ChangePasswd(pw *form.PasswordFrom) revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	pw.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Users.EditSettings)
	}

	user := CurrentUser(c.Controller)

	if err := user.UpdatePasswd(pw.NewPasswd, pw.CurrentPasswd); err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(Users.EditSettings)
	} else {
		c.Flash.Success("change password success")
	}

	return c.Redirect(Users.EditSettings)
}

func (c Users) Show(username string) revel.Result {
	user := model.FindUserByName(username)
	avatar := user.AvatarUrl()

	return c.Render(user, avatar)
}
