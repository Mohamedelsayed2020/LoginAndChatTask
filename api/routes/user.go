package routes

import (
	"LoginAndChatTask/api/controllers"
	"github.com/gorilla/mux"
	"net/http"
)
func Routes()  {

	r := mux.NewRouter()
	user := controllers.UserController{}
	r.HandleFunc("/user", user.CreateUser).Methods("POST")
	r.HandleFunc("/user", user.ListUsers).Methods("Get")
	r.HandleFunc("/user/{id}", user.GetUser).Methods("Get")
	r.HandleFunc("/user/{id}", user.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/login", user.Login).Methods("POST")


	http.ListenAndServe(":8000", r)
}