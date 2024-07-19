package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"OneGO/models"
	"OneGO/services"

	"github.com/gorilla/mux"
)

func GetMeasures(w http.ResponseWriter, r *http.Request) {
	measures := services.GetMeasures()
	json.NewEncoder(w).Encode(measures)
}

func GetMeasure(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	measure, found := services.GetMeasureByID(id)
	if !found {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(measure)
}

func CreateMeasure(w http.ResponseWriter, r *http.Request) {
	var measure models.Measure
	_ = json.NewDecoder(r.Body).Decode(&measure)
	services.CreateMeasure(&measure)
	json.NewEncoder(w).Encode(measure)
}

func UpdateMeasure(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var measure models.Measure
	_ = json.NewDecoder(r.Body).Decode(&measure)
	updatedMeasure, found := services.UpdateMeasure(id, &measure)
	if !found {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(updatedMeasure)
}

func DeleteMeasure(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	services.DeleteMeasure(id)
	w.WriteHeader(http.StatusNoContent)
}
