package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/braejan/platzi-go-rest-websockets/entities"
	"github.com/braejan/platzi-go-rest-websockets/server"
)

func HomeHandler(server server.Server) (handlerFunc http.HandlerFunc) {
	handlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		home := &entities.Home{
			Message: "this is a message from home.",
			Status:  true,
		}
		homeResponse := &entities.HomeResponse{
			Home:         home,
			ResponseCode: 200,
		}
		err := json.NewEncoder(w).Encode(homeResponse)
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
	return
}
