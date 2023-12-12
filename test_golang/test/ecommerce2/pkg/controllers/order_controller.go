package controllers

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"

	"test/ecommerce2/pkg/models"
	"test/ecommerce2/pkg/utils"
)

func (server *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	orderModel := models.Order{}
	utils.ParseBody(r, &orderModel)
	orderModel.BeforeCreate(server.DB)

	order, err := orderModel.CreateOrder(server.DB)
	if err != nil {
		res, _ := json.Marshal(&status{Check: "Sorry, Something wrong when create order"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}
	res, _ := json.Marshal(&order)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (server *Server) GetOrderByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		return
	}
	orderModel := models.Order{}
	_, err := orderModel.FindByUserID(server.DB, vars["id"])

	if err != nil {
		res, _ := json.Marshal(&status{Check: "Sorry, Something wrong when get order"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}
	res, _ := json.Marshal(&orderModel)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
