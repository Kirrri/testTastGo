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
	if products == nil {
		http.Error(w, "No products", http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(products)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	product, found := services.GetProductByID(id)
	if !found {
		http.Error(w, "There is no product with this ID", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	services.CreateProduct(&product)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	updatedProduct, found := services.UpdateProduct(id, &product)
	if !found {
		http.Error(w, "There is no product with this ID", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(updatedProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if !services.DeleteProduct(id) {
		http.Error(w, "There is no product with this ID", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
