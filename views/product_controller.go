package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"OneGO/models"
	"OneGO/services"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := services.GetProducts()
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	product, found := services.GetProductByID(id)
	if !found {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	services.CreateProduct(&product)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	updatedProduct, found := services.UpdateProduct(id, &product)
	if !found {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(updatedProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	services.DeleteProduct(id)
	w.WriteHeader(http.StatusNoContent)
}
