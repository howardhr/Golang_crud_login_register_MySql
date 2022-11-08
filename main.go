package main

import (
	"github.com/gorilla/mux"
	httplogger "github.com/jesseokeya/go-httplogger"
	"gomysql/commons"
	"gomysql/routes"
	"log"
	"net/http"
)

func main() {
	commons.Migrate()
	router := mux.NewRouter()
	routes.SetPersonasRoutes(router)
	routes.SetAuthRoutes(router)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	log.Println("Listening on..", "http://localhost:8080/")
	log.Fatal(server.ListenAndServe(), httplogger.Golog(router))
}
