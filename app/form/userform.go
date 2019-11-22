package form

import (
	"github.com/robfig/revel"
)

// user signup form field
type UserForm struct {
	Name           string
	Email          string
	Password       string
	RepeatPassword string
}

func (userform *UserForm) Validate(v *revel.Validation) {
	v.Required(userform.Name).Message("please verify you name")
	v.Required(userform.RepeatPassword).Message("please verify you repeat password ")
	v.Required(userform.Password).Message("please verify you password")
	v.Required(userform.RepeatPassword == userform.Password).Message("password do not match")
	v.Required(userform.Email).Message("please verify you Email")
	v.Email(userform.Email).Message("please verify you Email")
}

type UserLogin struct {
	Name     string
	Password string
	Remember bool
}

func (loginform *UserLogin) Validate(v *revel.Validation) {
	v.Required(loginform.Name)
	v.Required(loginform.Password)
}

// users settings form
type Settings struct {
	Site         string
	Weibo        string
	Introduction string
	Signature    string
	Github       string
}

// change password form

type PasswordFrom struct {
	CurrentPasswd string
	NewPasswd     string
	PasswdRepeat  string
}

// validates
func (pw *PasswordFrom) Validate(v *revel.Validation) {
	v.Required(pw.CurrentPasswd)
	v.Required(pw.NewPasswd)
	v.Required(pw.PasswdRepeat)
	v.Required(pw.PasswdRepeat == pw.NewPasswd).Message("you repeat password do not match")
}
