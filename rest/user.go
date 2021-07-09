package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"restbase/model"
	"strconv"

	"github.com/gorilla/mux"
)

func ViewAllUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	userList := model.Users.ViewAll(model.Users{})
	log.Printf("%+v", userList)
	json.NewEncoder(w).Encode(userList)

}
func SingleUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])
	json.NewEncoder(w).Encode(model.Users.View(model.Users{}, Id))

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	resBody, _ := ioutil.ReadAll(r.Body)
	var user model.Users
	json.Unmarshal(resBody, &user)
	user = model.Users.Add(user)
	log.Printf("Created:-\n%+v", user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	reqBody, _ := ioutil.ReadAll(r.Body)
	var User model.Users
	json.Unmarshal(reqBody, &User)
	User = model.Users.Update(User, id)
	log.Printf("Updated:-\n%+v", User)
	json.NewEncoder(w).Encode(User)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	model.Users.Delete(model.Users{}, id)
	log.Printf("Deleted User Details :- %+v", id)
	json.NewEncoder(w).Encode("SucessFully Deleted")
}

func Login(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("pass")

	var User model.Users
	User = model.Users.Login(User, email, password)
	log.Printf("Login By :-%+v", User)
	json.NewEncoder(w).Encode(User)
}
