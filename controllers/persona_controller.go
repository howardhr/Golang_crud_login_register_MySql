package controllers

import (
	json2 "encoding/json"
	"github.com/gorilla/mux"
	"gomysql/commons"
	"gomysql/models"
	"log"
	http "net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&personas)
	json, _ := json2.Marshal(personas)
	commons.SendResponse(w, http.StatusOK, json)
}

func Get(w http.ResponseWriter, r *http.Request) {
	persona := models.Persona{}
	db := commons.GetConnection()
	id := mux.Vars(r)["id"]

	defer db.Close()
	db.Find(&persona, id)
	if persona.ID > 0 {
		json, _ := json2.Marshal(persona)
		commons.SendResponse(w, http.StatusOK, json)
	} else {
		commons.SendError(w, http.StatusNotFound)
	}
}

func Save(w http.ResponseWriter, r *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()
	err := json2.NewDecoder(r.Body).Decode(&persona)

	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusBadRequest)
		return
	}
	err = db.Save(&persona).Error

	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusInternalServerError)
		return
	}
	json, _ := json2.Marshal(persona)
	commons.SendResponse(w, http.StatusCreated, json)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	persona := models.Persona{}
	db := commons.GetConnection()
	id := mux.Vars(r)["id"]

	defer db.Close()
	db.Find(&persona, id)
	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(w, http.StatusNotFound)
	}
}
