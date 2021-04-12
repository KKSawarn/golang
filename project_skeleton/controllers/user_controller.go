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
	resp, _ := user.UpdateUser(userData, id)
	if resp.ID == 0 {
		apihelpers.Respond(w, map[string]interface{}{"users": "user not exist"})
	} else {
		apihelpers.Respond(w, map[string]interface{}{"users": resp})
	}

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	resp := user.GetUsers()
	apihelpers.Respond(w, map[string]interface{}{"users": resp})

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	resp, _ := user.GetUser(id)
	if resp.ID == 0 {
		apihelpers.Respond(w, map[string]interface{}{"users": "user not found"})
	} else {
		apihelpers.Respond(w, map[string]interface{}{"users": resp})
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	resp, _ := user.DeleteUser(id)
	apihelpers.Respond(w, map[string]interface{}{"users": resp})
}
