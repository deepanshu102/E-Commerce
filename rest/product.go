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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	ResBody, _ := ioutil.ReadAll(r.Body)
	var Product model.Product
	json.Unmarshal(ResBody, &Product)
	Product = model.Product.Add(Product)
	json.NewEncoder(w).Encode(Product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	resBody, _ := ioutil.ReadAll(r.Body)
	var updatedProduct model.Product
	json.Unmarshal(resBody, &updatedProduct)
	updatedProduct = model.Product.Update(updatedProduct)
	log.Println(updatedProduct)
	json.NewEncoder(w).Encode(updatedProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])
	var DeletedProduct model.Product
	model.Product.Delete(DeletedProduct, Id)
	json.NewEncoder(w).Encode("message:Deleted Sucessfully")
}
func ViewProducts(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	listProduct := model.Product.ViewAll(model.Product{})
	log.Println("View products :- ", listProduct)
	json.NewEncoder(w).Encode(listProduct)
}

func SingleProduct(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	Vars := mux.Vars(r)
	Id, _ := strconv.Atoi(Vars["id"])
	var Product model.Product
	Product = model.Product.View(Product, Id)
	json.NewEncoder(w).Encode(Product)
}
