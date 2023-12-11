package controllers

import (
	"encoding/json"
	
	"net/http"

	// "strconv"

	"github.com/gorilla/mux"

	"test/ecommerce2/pkg/models"
	"test/ecommerce2/pkg/utils"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {
	productModel := models.Product{}
	products, err := productModel.GetProducts(server.DB)
	if err != nil {
		return
	}

	res, _ := json.Marshal(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (server *Server) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		return
	}

	productModel := models.Product{}
	product, err := productModel.FindByID(server.DB, vars["id"])
	if err != nil {
		return
	}

	res, _ := json.Marshal(&product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (server *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	productModel := models.Product{}
	utils.ParseBody(r, &productModel)
	err := productModel.CreateProduct(server.DB)
	if err != nil {
		res, _ := json.Marshal(&status{Check: "Sorry, Something wrong when create product"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}
	res, _ := json.Marshal(&status{Check: "create product OK"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (server *Server) DeleteProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	if vars["id"] == "" {
		return
	}

	productModel := models.Product{}
	product, err := productModel.DeleteProduct(server.DB, vars["id"])
	if err != nil {
		return
	}

	res, _ := json.Marshal(&product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
