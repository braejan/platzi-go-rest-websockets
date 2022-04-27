package main

import (
	"context"
	"log"
	"net/http"

	"github.com/braejan/platzi-go-rest-websockets/handlers"
	"github.com/braejan/platzi-go-rest-websockets/server"
	"github.com/braejan/platzi-go-rest-websockets/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	printErr(err, true)
	configEnvVariables := utils.GetConfigEnvVariables()
	configuration := &server.Config{
		Port:        configEnvVariables.Port,
		JWTSecret:   configEnvVariables.JWTSecret,
		DatabaseUrl: configEnvVariables.DatabaseUrl,
	}
	server, err := server.NewServer(context.Background(), configuration)
	printErr(err, true)
	server.Start(BindRoutes)

}

func BindRoutes(server server.Server, r *mux.Router) {
	r.Handle("/", handlers.HomeHandler(server)).Methods(http.MethodGet)
}

func printErr(err error, exit bool) {
	if err != nil {
		if exit {
			log.Fatalf("fatal error: %v", err)
		}
		log.Printf("error: %v", err)
	}
}
