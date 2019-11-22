package controllers

import (
	"github.com/robfig/revel"
	"webchat/app/model"
	//"fmt"
)

func isLogin(c *revel.Controller) bool {
	if _, ok := c.Session["user_name"]; ok {
		return true
	}
	return false
}

func (c Application) CheckUser() revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		//fmt.Println(c.Flash)
		return c.Redirect(Application.Index)
	}
	return nil
}

func (c Application) AddUser() revel.Result {
	if isLogin(c.Controller) {
		user := model.FindUserByName(c.Session["user_name"])
		c.RenderArgs["user"] = user
	}

	return nil
}

func CurrentUser(c *revel.Controller) (user *model.User) {
	if c.RenderArgs["user"] != nil {
		user = c.RenderArgs["user"].(*model.User)
		return
	}

	if isLogin(c) {
		user = model.FindUserByName(c.Session["user_name"])
		return
	}

	return nil
}
