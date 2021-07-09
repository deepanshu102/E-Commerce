package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-type", "application/json")
}

/*Category adding for the Rest Api*/
func CategoryHandler() {
	log.Print("category Rest Handler Added")
	myRouter.HandleFunc("/category", CategoryList).Methods("GET")
	myRouter.HandleFunc("/category", createCategory).Methods("POST")
	myRouter.HandleFunc("/category/{id}", DeleteCategory).Methods("DELETE")
	myRouter.HandleFunc("/category", UpdateCategory).Methods("PUT")
}

func ProductHandler() {
	log.Println("Product Handler added")
	myRouter.HandleFunc("/product/{id}", SingleProduct).Methods("GET")
	myRouter.HandleFunc("/product", ViewProducts).Methods("GET")
	myRouter.HandleFunc("/product", CreateProduct).Methods("POST")
	myRouter.HandleFunc("/product", UpdateProduct).Methods("PUT")
	myRouter.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")
}

func UserHandler() {
	log.Println("User Handler added")
	myRouter.HandleFunc("/userall", ViewAllUser).Methods("GET")
	myRouter.HandleFunc("/user/{id}", SingleUser).Methods("GET")
	myRouter.HandleFunc("/user", CreateUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/login", Login).Methods("POST")

}

func CartHandler() {
	log.Println("cart Handler added")
	myRouter.HandleFunc("/cart/{id}", ViewCarts).Methods("GET")
	myRouter.HandleFunc("/cart/{id}", CreateCart).Methods("POST")
	myRouter.HandleFunc("/cart/{id}", UpdateCart).Methods("PUT")
	myRouter.HandleFunc("/cart/{id}", DeleteCart).Methods("DELETE")

}

// func OrderHandler() {
// 	log.Println("Order Handler added")
// 	myRouter.HandleFunc("/order/{id}", ViewOrder).Methods("GET")
// 	myRouter.HandleFunc("/order/{id}", CreateOrder).Methods("POST")
// 	myRouter.HandleFunc("/order", UpdateOrder).Methods("PUT")
// }

func Handler() {
	fmt.Printf("\n----------REST HANDLER----------\n\n")
	CategoryHandler()
	ProductHandler()
	UserHandler()
	CartHandler()
	// OrderHandler()
	log.Println(http.ListenAndServe(":8080", myRouter))
}
