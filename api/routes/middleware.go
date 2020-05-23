package routes

import (
	"LoginAndChatTask/core/common"
	"LoginAndChatTask/core/model"
	"github.com/go-chi/chi"
	"net/http"
)

func IsLogged(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		sessionId := chi.URLParam(r, "SessionId")
		if sessionId == "" {
			common.Logger("error", sessionId)
			return
		}
		session := user.IsSessionExist(sessionId)
		if !session {
			common.Logger("error", "")
			return
		}
		f(w, r)
	}
}
