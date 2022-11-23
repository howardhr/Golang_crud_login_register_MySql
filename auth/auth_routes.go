package auth

import (
	"github.com/gorilla/mux"
)

func SetAuthRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/auth/api").Subrouter()
	subRoute.HandleFunc("/register", Register).Methods("POST")
	subRoute.HandleFunc("/login", Login).Methods("POST")
	subRoute.HandleFunc("/find/register", ValidateUserExist).Methods("GET")
}
