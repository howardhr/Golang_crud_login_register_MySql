package routes

import (
	"github.com/gorilla/mux"
	"gomysql/controllers"
)

func SetPersonasRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/user/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/add", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
}
