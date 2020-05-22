package controllers

import (
	"LoginAndChatTask/App"
	"LoginAndChatTask/common"
	"LoginAndChatTask/model"
	"LoginAndChatTask/server"
	"LoginAndChatTask/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
	App.Controller
}

func (self UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	user := model.NewUser()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	user.Password = common.HashPassword(user.Password)
	if err := server.Conn().Create(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (self UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	var user []model.User
	data := server.Conn().Find(&user)
	if data.Error != nil {
		json.NewEncoder(w).Encode(data.Error)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func (self UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	user := model.NewUser()
	result := server.Conn().Where("id = ?", id).First(&user)
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (self UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	user := model.NewUser()
	result := server.Conn().Where("id = ?", id).Find(&user)
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	result = server.Conn().Save(user)
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (self UserController) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var LoginRequest *services.UserLogin
	var User model.User

	if err := json.NewDecoder(r.Body).Decode(&LoginRequest); err != nil {
		self.JsonLogger(w, 500, "DecodingError", nil)
		return
	}

	_, user := LoginRequest.Format().ValidateLogin()

	userId := GetCurrentUserIdByEmail(LoginRequest.Email)

	errMsg, sessionId := model.CreateSession(userId)
	if errMsg != nil {
		self.JsonLogger(w, 500, "error while creating session"+errMsg.Error(), nil)
		return
	}

	user.SessionId = sessionId
	User.SessionId = user.SessionId

	// reset failed tries and update sessionId
	updateSessionId := user.UpdateUser(LoginRequest.Email)
	if updateSessionId != nil {
		self.JsonLogger(w, 400, "error while updating user Session", nil)
		return
	}

	self.Json(w, user, 200)
}

func GetCurrentUserIdByEmail(Email string) string {
	user := model.User{}
	if queryResult := server.Conn().Where(&model.User{Email: Email}).First(&user); queryResult.Error != nil {
		return queryResult.Error.Error()
	}
	return user.Id
}
