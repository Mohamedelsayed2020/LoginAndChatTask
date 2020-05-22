package routes

import (
	"LoginAndChatTask/controllers"
	"LoginAndChatTask/services"
	"html/template"
	"log"
	"net/http"
)


func homeHandler(tpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r)
	})
}

func Routes() {
	tpl := template.Must(template.ParseFiles("temp/home.html"))
	h := services.NewHub()
	router := http.NewServeMux()

	user := controllers.UserController{}
	router.HandleFunc("/user", user.CreateUser)
	router.HandleFunc("/user/login", user.Login)
	router.Handle("/", homeHandler(tpl))
	router.Handle("/ws/{sessionId}", services.WsHandler{H: h})
	log.Printf("serving on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}
