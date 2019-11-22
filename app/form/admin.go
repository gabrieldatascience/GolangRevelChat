package form

import (
	"github.com/robfig/revel"
)

type ServerSettings struct {
	Name string
}

func (sf *ServerSettings) Validate(v *revel.Validation) {
	v.Required(sf.Name)
}
