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
	if measures == nil {
		http.Error(w, "No measures", http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(measures)
	}
}

func GetMeasure(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	measure, found := services.GetMeasureByID(id)
	if !found {
		http.Error(w, "There is no measure with this ID", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(measure)
}

func CreateMeasure(w http.ResponseWriter, r *http.Request) {
	var measure models.Measure
	if err := json.NewDecoder(r.Body).Decode(&measure); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	services.CreateMeasure(&measure)
	json.NewEncoder(w).Encode(measure)
}

func UpdateMeasure(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var measure models.Measure
	if err := json.NewDecoder(r.Body).Decode(&measure); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	updatedMeasure, found := services.UpdateMeasure(id, &measure)
	if !found {
		http.Error(w, "There is no measure with this ID", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(updatedMeasure)
}

func DeleteMeasure(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if !services.DeleteMeasure(id) {
		http.Error(w, "There is no measure with this ID", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
