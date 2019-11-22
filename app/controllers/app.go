package controllers

import (
	"github.com/robfig/revel"
	//"fmt"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	return c.Render()
}
