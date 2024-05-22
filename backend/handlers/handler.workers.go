package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qemanuel/tech-sup-webapp/backend/models"
	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func GetWorkers(w http.ResponseWriter, r *http.Request) {
	table := persistence.DB.TablesMap["workers"]
	response, err := table.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		res, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetWorker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table := persistence.DB.TablesMap["workers"]
	response, err := table.Find(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	} else {
		res, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func CreateWorker(w http.ResponseWriter, r *http.Request) {
	// create Worker from request body
	var workerMap map[string]string
	json.NewDecoder(r.Body).Decode(&workerMap)
	worker, err := models.NewWorker(workerMap["name"], workerMap["email"], workerMap["phone"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		// add worker to DB
		table := persistence.DB.TablesMap["workers"]
		response, err := table.Add(worker)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			res, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		}
	}
}

func UpdateWorker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// create Worker from request body
	var workerMap map[string]string
	json.NewDecoder(r.Body).Decode(&workerMap)
	worker, err := models.NewWorker(workerMap["name"], workerMap["email"], workerMap["phone"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		// Update worker on DB
		table := persistence.DB.TablesMap["workers"]
		id := vars["id"]
		response, err := table.Update(worker, id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		res, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func DeleteWorker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table := persistence.DB.TablesMap["workers"]
	err := table.Remove(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Worker not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Worker deleted"))
	}
}
