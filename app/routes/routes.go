// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tRoomApi struct {}
var RoomApi tRoomApi


func (p tRoomApi) Users(
		roomkey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomkey", roomkey)
	return revel.MainRouter.Reverse("RoomApi.Users", args).Url
}


type tApplication struct {}
var Application tApplication


func (p tApplication) CheckUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.CheckUser", args).Url
}

func (p tApplication) AddUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.AddUser", args).Url
}

func (p tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).Url
}


type tWebsocket struct {}
var Websocket tWebsocket


func (p tWebsocket) Chat(
		roomkey string,
		ws interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomkey", roomkey)
	revel.Unbind(args, "ws", ws)
	return revel.MainRouter.Reverse("Websocket.Chat", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (p tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (p tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (p tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (p tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (p tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tUsers struct {}
var Users tUsers


func (p tUsers) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.New", args).Url
}

func (p tUsers) Create(
		userform interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userform", userform)
	return revel.MainRouter.Reverse("Users.Create", args).Url
}

func (p tUsers) MyRooms(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.MyRooms", args).Url
}

func (p tUsers) EditSettings(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.EditSettings", args).Url
}

func (p tUsers) SaveSettings(
		setting interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "setting", setting)
	return revel.MainRouter.Reverse("Users.SaveSettings", args).Url
}

func (p tUsers) ChangePasswd(
		pw interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "pw", pw)
	return revel.MainRouter.Reverse("Users.ChangePasswd", args).Url
}

func (p tUsers) Show(
		username string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	return revel.MainRouter.Reverse("Users.Show", args).Url
}


type tRooms struct {}
var Rooms tRooms


func (p tRooms) Index(
		p int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "p", p)
	return revel.MainRouter.Reverse("Rooms.Index", args).Url
}

func (p tRooms) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Rooms.New", args).Url
}

func (p tRooms) Create(
		rf interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "rf", rf)
	return revel.MainRouter.Reverse("Rooms.Create", args).Url
}

func (p tRooms) Show(
		roomkey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomkey", roomkey)
	return revel.MainRouter.Reverse("Rooms.Show", args).Url
}

func (p tRooms) Edit(
		roomkey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomkey", roomkey)
	return revel.MainRouter.Reverse("Rooms.Edit", args).Url
}

func (p tRooms) Update(
		roomkey string,
		updateroom interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomkey", roomkey)
	revel.Unbind(args, "updateroom", updateroom)
	return revel.MainRouter.Reverse("Rooms.Update", args).Url
}

func (p tRooms) Logs(
		roomkey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomkey", roomkey)
	return revel.MainRouter.Reverse("Rooms.Logs", args).Url
}


type tSessions struct {}
var Sessions tSessions


func (p tSessions) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sessions.New", args).Url
}

func (p tSessions) Create(
		loginform interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "loginform", loginform)
	return revel.MainRouter.Reverse("Sessions.Create", args).Url
}

func (p tSessions) Destroy(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sessions.Destroy", args).Url
}


type tFavorite struct {}
var Favorite tFavorite


func (p tFavorite) FavoriteRoom(
		roomkey string,
		like bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomkey", roomkey)
	revel.Unbind(args, "like", like)
	return revel.MainRouter.Reverse("Favorite.FavoriteRoom", args).Url
}


type tAdmin struct {}
var Admin tAdmin


func (p tAdmin) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Index", args).Url
}

func (p tAdmin) Users(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Users", args).Url
}

func (p tAdmin) Rooms(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Rooms", args).Url
}

func (p tAdmin) ChangeLogStatus(
		roomKey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomKey", roomKey)
	return revel.MainRouter.Reverse("Admin.ChangeLogStatus", args).Url
}

func (p tAdmin) ChangeRunStatus(
		roomKey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomKey", roomKey)
	return revel.MainRouter.Reverse("Admin.ChangeRunStatus", args).Url
}

func (p tAdmin) SiteSettings(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.SiteSettings", args).Url
}

func (p tAdmin) SaveServerSettings(
		sf interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "sf", sf)
	return revel.MainRouter.Reverse("Admin.SaveServerSettings", args).Url
}

func (p tAdmin) OnlineUsers(
		roomKey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomKey", roomKey)
	return revel.MainRouter.Reverse("Admin.OnlineUsers", args).Url
}


