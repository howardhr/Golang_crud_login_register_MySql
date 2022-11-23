package auth

import (
	json2 "encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gomysql/commons"
	"gomysql/user/models"
	"log"
	"net/http"
)

func ValidateUserExist(w http.ResponseWriter, r *http.Request) {
	registers := []models.Usuario{}
	db := commons.GetConnection()
	defer db.Close()
	db.Find(&registers)
	json, _ := json2.Marshal(registers)
	commons.SendResponse(w, http.StatusOK, json)
}

func HashContrasena(Contrasena string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Contrasena), 14)
	return string(bytes), err
}

func CheckContrasenaHash(Contrasena, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Contrasena))
	return err == nil
}

/**
 * Metodo para registrar un usuario
 * @param {http.ResponseWriter} singUP registered
 * @param {*http.Request}
 * @return {singUP<Object>}
 */
func Register(w http.ResponseWriter, r *http.Request) {

	singUP := models.Usuario{}
	registered := models.Usuario{}
	hash, _ := HashContrasena(singUP.Contrasena)
	db := commons.GetConnection()
	defer db.Close()
	err := json2.NewDecoder(r.Body).Decode(&singUP)
	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusBadRequest)
		return
	}

	db.Find(&registered)
	fmt.Println("register", registered)
	fmt.Println("login", singUP)
	json, _ := json2.Marshal(registered)
	json, _ = json2.Marshal(singUP)
	json, _ = json2.Marshal(hash)
	fmt.Println("Hash:    ", hash)
	if registered.Documento == singUP.Documento {
		commons.SendError(w, http.StatusBadRequest)
		json, _ = json2.Marshal("Ya existe un usuario creado con este numero de documento")
	} else {
		singUP.Contrasena = hash
		err = db.Save(&singUP).Error

		json, _ = json2.Marshal("Usuario creado correctamente")
		commons.SendResponse(w, http.StatusOK, []byte(`{}`))
	}

	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusInternalServerError)
		return
	}
	commons.SendResponse(w, http.StatusCreated, json)
}

/**
 * Metodo para iniciarsesion
 * @param {http.ResponseWriter} singUP registered
 * @param {*http.Request}
 * @return {singUP<Object>}
 */
func Login(w http.ResponseWriter, r *http.Request) {
	login := models.Usuario{}
	register := models.Usuario{}
	db := commons.GetConnection()
	defer db.Close()
	err := json2.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusBadRequest)
		return
	}
	db.Find(&register)

	json, _ := json2.Marshal(register)
	json, _ = json2.Marshal(login.Contrasena)

	fmt.Println("passw:    ", register.Contrasena)
	fmt.Println("passwLogin:    ", login.Contrasena)
	loginBody := login.Contrasena
	registerHash := register.Contrasena
	//err = bcrypt.CompareHashAndPassword(loginBody, registerHash)
	match := CheckContrasenaHash(loginBody, registerHash)
	fmt.Println("err:   ", match)
	if match != false {
		json, _ = json2.Marshal("iniciaste session")
		log.Println("iniciaste session")
	} else {
		json, _ = json2.Marshal("Datos incorrectos")
		log.Println("Datos incorrectos")
	}

	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusInternalServerError)
		return
	}
	commons.SendResponse(w, http.StatusCreated, json)

}
