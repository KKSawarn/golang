package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"usercurd/apihelpers"
	"usercurd/models"
	"usercurd/services/user"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var userData models.User
	json.NewDecoder(r.Body).Decode(&userData)
	resp := user.CreateUser(userData)
	apihelpers.Respond(w, map[string]interface{}{"users": resp})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var userData models.User
	id, _ := strconv.Atoi(params["id"])
	json.NewDecoder(r.Body).Decode(&userData)
	resp := user.UpdateUser(userData, id)
	apihelpers.Respond(w, map[string]interface{}{"users": resp})

}
func GetUsers(w http.ResponseWriter, r *http.Request) {

	//call service
	resp := user.GetUsers()
	//return response using api helper
	apihelpers.Respond(w, map[string]interface{}{"users": resp})

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	//call service
	resp := user.GetUser(id)
	//return response using api helper
	apihelpers.Respond(w, map[string]interface{}{"users": resp})

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	resp := user.DeleteUser(id)
	apihelpers.Respond(w, map[string]interface{}{"users": resp})
}
