package handlers

import (
	"../../data/database"
	"../../data/models"
	"../../decoder"
	"../../logger"
	"encoding/json"
	"net/http"
)

func AddEquipment(w http.ResponseWriter, r *http.Request) {
	var equipment models.Equipment
	decoder.Get(r.Body, &equipment)

	err := database.AddEquipment(equipment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(Response{Errors: Errors{Global: "Failed to create equipment:\n" + err.Error()}}); err != nil {
			logger.LogErr(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetEquipments(w http.ResponseWriter, r *http.Request) {
	equipments, err := database.GetEquipments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(Response{Errors: Errors{Global: "Failed to get equipments:\n" + err.Error()}}); err != nil {
			logger.LogErr(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Equipments: equipments}); err != nil {
		logger.LogErr(err)
	}
}

func GetEquipment(w http.ResponseWriter, r *http.Request) {
	var equipment models.Equipment
	decoder.Get(r.Body, &equipment)
	temp, err := database.GetEquipment(equipment.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(Response{Errors: Errors{Global: "Failed to get equipment:\n" + err.Error()}}); err != nil {
			logger.LogErr(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Equipments: []models.Equipment{temp}}); err != nil {
		logger.LogErr(err)
	}
}