package controllers

import (
	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	// server.Router.HandleFunc("/", server.Home).Methods("GET")

	server.Router.HandleFunc("/login", server.DoLogin).Methods("POST")
	server.Router.HandleFunc("/register", server.DoRegister).Methods("POST")
	server.Router.HandleFunc("/logout", server.Logout).Methods("GET")
	server.Router.HandleFunc("/user/{id}", server.DeleteUser).Methods("DELETE")
	server.Router.HandleFunc("/user/{id}", server.GetUser).Methods("GET")
	server.Router.HandleFunc("/user/{id}", server.UpdateUser).Methods("PUT")

	server.Router.HandleFunc("/products", server.Products).Methods("GET")
	server.Router.HandleFunc("/products/{id}", server.GetProductByID).Methods("GET")
	server.Router.HandleFunc("/products", server.CreateProduct).Methods("POST")
	server.Router.HandleFunc("/products/{id}", server.UpdateProduct).Methods("PUT")
	server.Router.HandleFunc("/products/{id}", server.DeleteProduct).Methods("DELETE")

	server.Router.HandleFunc("/category", server.GetCategories).Methods("GET")
	server.Router.HandleFunc("/category", server.CreateCategory).Methods("POST")
	server.Router.HandleFunc("/category/{id}", server.GetCategoryByID).Methods("GET")
	server.Router.HandleFunc("/category/{id}", server.DeleteCategory).Methods("DELETE")
	server.Router.HandleFunc("/category/{id}", server.UpdateCategory).Methods("PUT")

	server.Router.HandleFunc("/product-img/{id}", server.CreateImageByProductID).Methods("POST")
	server.Router.HandleFunc("/product-img/{id}", server.DeleteImageByProductID).Methods("DELETE")

	server.Router.HandleFunc("/carts/{id}", server.GetCart).Methods("GET")
	server.Router.HandleFunc("/carts", server.AddItem).Methods("POST")
	server.Router.HandleFunc("/carts/{id}", server.DeleteItem).Methods("DELETE")

	server.Router.HandleFunc("/orders", server.CreateOrder).Methods("POST")
	server.Router.HandleFunc("/orders/{id}", server.GetOrderByUserID).Methods("GET")
}
