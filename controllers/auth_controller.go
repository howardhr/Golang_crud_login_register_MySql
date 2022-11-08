package controllers

import (
	json2 "encoding/json"
	"gomysql/commons"
	"gomysql/models"
	"log"
	"net/http"
)

func ValidateUserExist(w http.ResponseWriter, r *http.Request) {
	registers := []models.Register{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&registers)
	json, _ := json2.Marshal(registers)
	commons.SendResponse(w, http.StatusOK, json)
}

func Register(w http.ResponseWriter, r *http.Request) {
	register := models.Register{}
	registers := models.Register{}
	db := commons.GetConnection()
	defer db.Close()
	err := json2.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusBadRequest)
		return
	}

	db.Find(&registers)
	json, _ := json2.Marshal(registers)
	if registers.Documento == register.Documento {
		log.Println("ya este usuario existe")
	} else {

		err = db.Save(&register).Error
	}

	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusInternalServerError)
		return
	}

	commons.SendResponse(w, http.StatusCreated, json)
}

func Login(w http.ResponseWriter, r *http.Request) {
	login := models.Register{}
	registers := models.Register{}
	db := commons.GetConnection()
	defer db.Close()
	err := json2.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusBadRequest)
		return
	}

	db.Find(&registers)
	json, _ := json2.Marshal(registers)
	json, _ = json2.Marshal(login)
	if login == registers {
		log.Println("iniciaste session")
	} else {
		log.Println("Datos incorrectos")
	}

	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusInternalServerError)
		return
	}

	commons.SendResponse(w, http.StatusCreated, json)
}
