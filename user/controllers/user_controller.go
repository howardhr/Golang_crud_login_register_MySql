package controllers

import (
	json2 "encoding/json"
	"github.com/gorilla/mux"
	"gomysql/commons"
	"gomysql/user/models"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	user := []models.Usuario{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&user)
	json, _ := json2.Marshal(user)
	commons.SendResponse(w, http.StatusOK, json)
}

func Get(w http.ResponseWriter, r *http.Request) {
	user := models.Usuario{}
	db := commons.GetConnection()
	id := mux.Vars(r)["id"]

	defer db.Close()
	db.Find(&user, id)
	if user.ID > 0 {
		json, _ := json2.Marshal(user)
		commons.SendResponse(w, http.StatusOK, json)
	} else {
		commons.SendError(w, http.StatusNotFound)
	}
}

/*func Save(w http.ResponseWriter, r *http.Request) {
	user := models.Usuario{}

	db := commons.GetConnection()
	defer db.Close()
	err := json2.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusBadRequest)
		return
	}
	err = db.Save(&user).Error
	log.Println("Usuario creado con exito")
	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusInternalServerError)
		return
	}
	json, _ := json2.Marshal(user)
	commons.SendResponse(w, http.StatusCreated, json)
}*/

func Delete(w http.ResponseWriter, r *http.Request) {
	user := models.Usuario{}
	db := commons.GetConnection()
	id := mux.Vars(r)["id"]

	defer db.Close()
	db.Find(&user, id)
	if user.ID > 0 {
		db.Delete(user)
		commons.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(w, http.StatusNotFound)
	}
}
