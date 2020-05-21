package controllers

import (
	"LoginAndChatTask/api/App"
	"LoginAndChatTask/api/services"
	"LoginAndChatTask/model"
	"LoginAndChatTask/model/common"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
	App.Controller
}

func (self UserController)CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	user := model.NewUser()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	user.Password = common.HashPassword(user.Password)
	if err := common.Conn().Create(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (self UserController)ListUsers(w http.ResponseWriter, r *http.Request)  {
	var user []model.User
	data := common.Conn().Find(&user)
	if data.Error != nil {
		json.NewEncoder(w).Encode(data.Error)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func (self UserController)GetUser(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	id := params["id"]
	user := model.NewUser()
	result := common.Conn().Where("id = ?", id).First(&user)
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (self UserController)UpdateUser(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	id := params["id"]
	user := model.NewUser()
	result := common.Conn().Where("id = ?", id).Find(&user)
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	result = common.Conn().Save(user)
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func  (self UserController)Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var userLogin *services.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		self.JsonLogger(w, 500, "DecodingError", err)
		self.Logger("DecodingError", "error")
		return
	}
	err, user := userLogin.Format().ValidateLogin()
	if err != "" {
		self.Json(w, err, 200)
		return
	}
	self.Json(w, user, 200)

}