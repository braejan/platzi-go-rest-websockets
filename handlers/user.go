package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/braejan/platzi-go-rest-websockets/server"
	"github.com/braejan/platzi-go-rest-websockets/usecase"
)

type SingUpRequest struct {
	Email string `json:"email"`
}

type SingUpResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func SingUpHandler(server server.Server) (handler http.HandlerFunc) {
	handler = func(w http.ResponseWriter, r *http.Request) {
		var request = &SingUpRequest{}
		err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			http.Error(w, http.ErrBodyNotAllowed.Error(), http.StatusBadRequest)
			return
		}
		context := r.Context()
		usecases, err := usecase.NewUsecases(&context)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = usecases.User.CreateUser(&context, request.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	return
}
