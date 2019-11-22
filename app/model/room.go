package model

import (
	//"github.com/robfig/revel"
	"errors"
	"github.com/hoisie/redis"
	"log"
	"strconv"
	"strings"
	"time"
	"webchat/app/form"
)

const (
	PageSize int = 12
)

var redisClient redis.Client

type Room struct {
	Id          int `pk`
	UserId      int
	RoomKey     string
	Title       string
	Private     bool
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func AllRoom() []Room {
	var rooms []Room
	db := GetDblink()
	db.FindAll(&rooms)
	return rooms
}

func RoomCount() int {
	db := GetDblink()
	var itemCount int

	err := db.Db.QueryRow("select count(*) as count from room").Scan(&itemCount)

	if err != nil {
		panic(err)
	}

	return itemCount
}

func FindRoomByUserId(user_id int) []Room {
	db := GetDblink()
	var rooms []Room
	log.Println("userid is--:", user_id)

	db.Where("user_id=?", user_id).FindAll(&rooms)
	return rooms
}

func FindOnePage(p int) []Room {
	db := GetDblink()

	var rooms []Room
	var offset int

	if p == 0 {
		offset = 0
	} else {
		offset = (p - 1) * PageSize
	}

	db.Limit(PageSize, offset).FindAll(&rooms)
	return rooms
}

func FindRoomByRoomKey(rk string) *Room {
	db := GetDblink()
	var room Room

	if err := db.Where("room_key=?", rk).Find(&room); err != nil {
		return nil
	}

	return &room
}

func NewRoom(form *form.RoomForm) (room *Room) {
	room = &Room{
		UserId:      form.UserId,
		RoomKey:     form.RoomKey,
		Title:       form.Title,
		Private:     form.Private,
		Description: form.Desc,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return room
}

func (room *Room) Save() (*Room, error) {
	db := GetDblink()

	if err := room.ValidatesUniqueness(); err != nil {
		return nil, err
	}

	if err := db.Save(room); err != nil {
		return nil, err
	}

	return room, nil
}

func (room *Room) Update(form *form.UpdateRoom) error {
	db := GetDblink()

	room.Description = form.Desc
	room.Title = form.Title

	if err := db.Save(room); err != nil {
		return err
	}

	return nil
}

func (room *Room) ValidatesUniqueness() error {
	db := GetDblink()
	var r Room
	db.Where("room_key=?", room.RoomKey).Find(&r)
	if r.Id != 0 {
		return errors.New("input room id: " + room.RoomKey + " has exist")
	}
	return nil
}

// room index recent user

type RoomList struct {
	Id          int
	RoomKey     string
	Description string
	Title       string
	Private     bool
	RecentUsers []*RecentUser
}

type RecentUser struct {
	Id     int
	Avatar string
	Name   string
}

func (r *Room) GetRecentUsers() []*RecentUser {
	userIds, _ := redisClient.Smembers("room:" + r.RoomKey + ":users")

	if r.RoomKey == "aaa" {
		log.Println(len(userIds))
	}
	var recentusers []*RecentUser

	for _, id := range userIds {
		key := "user:" + string(id)
		if r.RoomKey == "aaa" {
			log.Println(key)
		}

		user := make(map[string]string)
		err := redisClient.Hgetall(key, user)

		if err == nil {
			userid, _ := strconv.Atoi(user["id"])

			ru := &RecentUser{
				Id:     userid, //to int
				Avatar: user["avatar"],
				Name:   user["name"],
			}

			recentusers = append(recentusers, ru)
		}
	}

	if len(recentusers) > 5 {
		recentusers = recentusers[:5]
	}

	return recentusers
}

// last message from redis
type LatestMessage struct {
	Type     string
	UserName string
	Text     string
	Time     string
}

func AllMessageFromRedis(roomkey string) (LM []*LatestMessage) {
	LM = GetMessageFromRedis(roomkey, 0, -1)
	return
}

func (r *Room) LatestMessage() (LM []*LatestMessage) {
	LM = GetMessageFromRedis(r.RoomKey, 0, 4)
	return
}

func GetMessageFromRedis(roomkey string, start int, end int) (LM []*LatestMessage) {
	key := "room:" + roomkey
	messages, _ := redisClient.Lrange(key, start, end)

	for _, m := range messages {
		ms := strings.Split(string(m), "|")

		latestMessage := &LatestMessage{
			Type:     ms[0],
			UserName: ms[1],
			Text:     ms[2],
			Time:     ms[3],
		}

		LM = append(LM, latestMessage)
	}

	return LM

}
