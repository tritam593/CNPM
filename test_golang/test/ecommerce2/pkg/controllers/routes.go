package controllers

import (
	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	// server.Router.HandleFunc("/", server.Home).Methods("GET")

	// server.Router.HandleFunc("/login", server.Login).Methods("GET")
	server.Router.HandleFunc("/login", server.DoLogin).Methods("POST")
	// server.Router.HandleFunc("/register", server.Register).Methods("GET")
	server.Router.HandleFunc("/register", server.DoRegister).Methods("POST")
	server.Router.HandleFunc("/logout", server.Logout).Methods("GET")
	server.Router.HandleFunc("/remove-user", server.DeleteUser).Methods("DELETE")

	server.Router.HandleFunc("/products", server.Products).Methods("GET")
	server.Router.HandleFunc("/products/{id}", server.GetProductByID).Methods("GET")
	server.Router.HandleFunc("/products", server.CreateProduct).Methods("POST")
	server.Router.HandleFunc("/products/{id}", server.DeleteProduct).Methods("DELETE")

	server.Router.HandleFunc("/category", server.GetCategories).Methods("GET")
	server.Router.HandleFunc("/category", server.CreateCategory).Methods("POST")
	server.Router.HandleFunc("/category/{id}", server.GetCategoryByID).Methods("GET")
	server.Router.HandleFunc("/category/{id}", server.DeleteCategory).Methods("DELETE")
	server.Router.HandleFunc("/category/{id}", server.UpdateCategory).Methods("PUT")

	// server.Router.HandleFunc("/carts", server.GetCart).Methods("GET")
	// server.Router.HandleFunc("/carts", server.AddItemToCart).Methods("POST")
	// server.Router.HandleFunc("/carts/update", server.UpdateCart).Methods("POST")
	// server.Router.HandleFunc("/carts/cities", server.GetCitiesByProvince).Methods("GET")
	// server.Router.HandleFunc("/carts/calculate-shipping", server.CalculateShipping).Methods("POST")
	// server.Router.HandleFunc("/carts/apply-shipping", server.ApplyShipping).Methods("POST")
	// server.Router.HandleFunc("/carts/remove/{id}", server.RemoveItemByID).Methods("GET")

	// server.Router.HandleFunc("/orders/checkout", server.Checkout).Methods("POST")
	// server.Router.HandleFunc("/orders/{id}", server.ShowOrder).Methods("GET")
}
