package controllers

import (
	"github.com/robfig/revel"
	"log"
	//"webchat/app/form"
	//"webchat/app/model"
)

type Favorite struct {
	*Application
}

func (c Favorite) FavoriteRoom(roomkey string, like bool) revel.Result {
	log.Println("roomke is", roomkey)
	log.Println("favorites request body is", c.Request.Body)
	return nil
}
