package routes

import (
	"LoginAndChatTask/api/controllers"
	"LoginAndChatTask/model/common"
	"github.com/gorilla/mux"
	"net/http"
)
func Routes()  {

	r := mux.NewRouter()
	sr := http.FileServer(http.Dir("../temp"))

	user := controllers.UserController{}
	r.HandleFunc("/user", user.CreateUser).Methods("POST")
	r.HandleFunc("/user", user.ListUsers).Methods("Get")
	r.HandleFunc("/user/{id}", user.GetUser).Methods("Get")
	r.HandleFunc("/user/{id}", user.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/login", user.Login).Methods("POST")
	http.Handle("/", sr)
	http.HandleFunc("/ws", common.HandleConnections)

	http.ListenAndServe(":8080", r)
}