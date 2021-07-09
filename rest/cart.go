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

func CreateCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	ResBody, _ := ioutil.ReadAll(r.Body)
	var cartList []model.Cart
	var cart model.Cart
	json.Unmarshal(ResBody, &cart)
	cartList = model.Cart.Add(cart, userId)
	json.NewEncoder(w).Encode(cartList)
}

func UpdateCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	resBody, _ := ioutil.ReadAll(r.Body)
	var updatedCart []model.Cart
	var cart model.Cart
	json.Unmarshal(resBody, &cart)
	updatedCart = model.Cart.Update(cart)
	log.Println(updatedCart)
	json.NewEncoder(w).Encode(updatedCart)
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])
	var DeletedCart model.Cart
	model.Cart.Delete(DeletedCart, Id)
	json.NewEncoder(w).Encode("Sucessfully Deleted")
}
func ViewCarts(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])
	listCart := model.Cart.View(model.Cart{}, Id)
	log.Println("View Carts :- ", listCart)
	json.NewEncoder(w).Encode(listCart)
}
