package routes

import (
	"github.com/gorilla/mux"
	"gomysql/controllers"
)

func SetAuthRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/auth/api").Subrouter()
	subRoute.HandleFunc("/register", controllers.Register).Methods("POST")
	subRoute.HandleFunc("/login", controllers.Login).Methods("POST")
	subRoute.HandleFunc("/find/register", controllers.ValidateUserExist).Methods("GET")
}
