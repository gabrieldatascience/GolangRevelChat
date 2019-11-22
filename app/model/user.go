package model

import (
	//"github.com/robfig/revel"
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
	"webchat/app/form"
)

type User struct {
	Id            int `pk`
	Name          string
	Email         string
	Salt          string
	Encryptpasswd string
	Site          string
	Weibo         string
	Introduction  string
	Signature     string
	// Avatar        string
	Github  string
	Created time.Time
	Updated time.Time
}

// generate a User form input form field
func NewUser(form *form.UserForm) (user *User) {
	user = new(User)
	user.Id = 0
	user.Name = form.Name
	user.Email = form.Email
	user.Salt = generate_salt()
	user.Encryptpasswd = encryptPassword(form.Password, user.Salt)
	user.Created, user.Updated = time.Now(), time.Now()

	return user
}

func (user *User) ValidatesUniqueness() error {
	db := GetDblink()
	var u User
	db.Where("name=?", user.Name).Find(&u)

	if u.Id != 0 {
		return errors.New("input name: " + user.Name + " has exist")
	}

	db.Where("email=?", user.Email).Find(&u)
	if u.Id != 0 {
		return errors.New("input email: " + user.Email + " has exist")
	}

	return nil
}

func (user *User) Save() error {
	db := GetDblink()

	if err := user.ValidatesUniqueness(); err != nil {
		return err
	}

	if err := db.Save(user); err != nil {
		return err
	}

	return nil
}

func FindUserByName(name string) *User {
	db := GetDblink()
	user := new(User)

	if err := db.Where("name=?", name).Find(user); err != nil {
		panic(err)
	}
	return user
}

// auth user
func Authenticate(name string, password string) bool {
	db := GetDblink()
	var user User
	err := db.Where("name=?", name).Find(&user)

	if err != nil {
		return false
	}

	if encryptPassword(password, user.Salt) == user.Encryptpasswd {
		return true
	}
	return false
}

// for generate rand salt
func generate_salt() string {
	rand.Seed(time.Now().UnixNano())
	salt := fmt.Sprintf("%x", rand.Int31())
	return salt
}

// for enrypt password
func encryptPassword(password, salt string) string {
	h := sha1.New()
	h.Write([]byte(password + salt))
	bs := fmt.Sprintf("%x", h.Sum(nil))
	return bs
}

// get a user's rooms
func (u *User) Rooms() []Room {
	var rooms []Room
	rooms = FindRoomByUserId(u.Id)
	return rooms
}

func (u *User) SaveSettings(setting *form.Settings) error {
	db := GetDblink()

	u.Weibo = setting.Weibo
	u.Site = setting.Site
	u.Introduction = setting.Introduction
	u.Signature = setting.Signature
	u.Github = setting.Github

	if err := db.Save(u); err != nil {
		return err
	}

	return nil
}

func (u *User) UpdatePasswd(newPasswd, currentPasswd string) error {
	db := GetDblink()

	if !Authenticate(u.Name, currentPasswd) {
		return errors.New("you  curent password is error!")
	} else {
		u.Encryptpasswd = encryptPassword(newPasswd, u.Salt)
		if err := db.Save(u); err != nil {
			return err
		}
	}

	return nil
}

func Hash(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	hash := md5.New()
	hash.Write([]byte(email))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (u *User) AvatarUrl() string {
	// if u.Avatar == "" {
	// 	return "/public/avatar/" + size + "_default.png"
	// }

	// return u.Avatar
	log.Println(u.Email)
	return "http://www.gravatar.com/avatar/" + Hash(u.Email)
}

func LatestUsers(count int) (users []User) {
	db := GetDblink()
	db.Limit(count).FindAll(&users)
	return
}

func AllUsers() (users []User) {
	db := GetDblink()
	db.FindAll(&users)
	return
}

func UserCount() int {
	db := GetDblink()
	var itemCount int

	err := db.Db.QueryRow("select count(*) as count from user").Scan(&itemCount)

	if err != nil {
		panic(err)
	}

	return itemCount
}
