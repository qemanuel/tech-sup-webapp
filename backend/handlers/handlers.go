package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qemanuel/tech-sup-webapp/backend/models"
	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func GetWorkers(w http.ResponseWriter, r *http.Request) {
	response, err := persistence.GetAll("workers")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(response)
}

func GetWorker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := persistence.Find("workers", vars["id"])
	res, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateWorker(w http.ResponseWriter, r *http.Request) {
	var worker models.Worker
	json.NewDecoder(r.Body).Decode(&worker)
	id, err := persistence.Add("workers", worker)
	idString := fmt.Sprint(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	worker.SetId(idString)
	json.NewEncoder(w).Encode(&worker)
}

func UpdateWorker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var worker models.Worker
	json.NewDecoder(r.Body).Decode(&worker)
	err := persistence.Update("workers", worker, vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&worker)
}

func DeleteWorker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := persistence.Remove("workers", vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Worker not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
}
