package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"test/e-commerce-test/pkg/models"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	perPage := 9

	productModel := models.Product{}
	products, totalRows, err := productModel.GetProducts(server.DB, perPage, page)
	if err != nil {
		return
	}

	pagination, _ := GetPaginationLinks(server.AppConfig, PaginationParams{
		Path:        "products",
		TotalRows:   int32(totalRows),
		PerPage:     int32(perPage),
		CurrentPage: int32(page),
	})

	temp := map[string]interface{}{
		"products":   products,
		"pagination": pagination,
		"user":       server.CurrentUser(w, r),
	}

	res, _ := json.Marshal(&temp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (server *Server) GetProductBySlug(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	if vars["slug"] == "" {
		return
	}

	productModel := models.Product{}
	product, err := productModel.FindBySlug(server.DB, vars["slug"])
	if err != nil {
		return
	}

	temp := map[string]interface{}{
		"product": product,
		"success": GetFlash(w, r, "success"),
		"error":   GetFlash(w, r, "error"),
		"user":    server.CurrentUser(w, r),
	}

	res, _ := json.Marshal(&temp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
