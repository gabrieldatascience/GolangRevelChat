package form

import (
	"github.com/robfig/revel"
)

type RoomForm struct {
	UserId  int
	RoomKey string
	Title   string
	Desc    string
	Private bool
}

func (rf *RoomForm) Validate(v *revel.Validation) {
	v.Required(rf.UserId)
	v.Required(rf.RoomKey)
	v.Required(rf.Title)
	v.Required(rf.Desc)
}

type UpdateRoom struct {
	Title string
	Desc  string
}

func (upf *UpdateRoom) Validate(v *revel.Validation) {
	v.Required(upf.Title)
	v.Required(upf.Desc)
}
