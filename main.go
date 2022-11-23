package main

import (
	"github.com/go-chi/httplog"
	"github.com/gorilla/mux"
	"gomysql/auth"
	"gomysql/commons"
	routes2 "gomysql/user/routes"
	"log"
	"net/http"
)

func main() {
	// Logger
	logger := httplog.NewLogger("httplog", httplog.Options{
		Concise: true,
	})
	commons.Migrate()
	router := mux.NewRouter()
	routes2.SetPersonasRoutes(router)
	auth.SetAuthRoutes(router)
	router.Use(httplog.RequestLogger(logger))
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	log.Println("Listening on..", "http://localhost:8080/")
	log.Fatal(server.ListenAndServe())
}
