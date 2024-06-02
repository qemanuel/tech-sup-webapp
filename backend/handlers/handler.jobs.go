package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qemanuel/tech-sup-webapp/backend/models"
	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {
	searchKeys := []string{
		"id",
		"status",
		"device_id",
		"author_id",
		"assigned_id",
		//		"customer_id",
		//		"before",
		//		"after",
	}
	urlQuery := r.URL.Query()
	queryMap := make(map[string]string, len(urlQuery))
	for _, key := range searchKeys {
		queryValue := r.URL.Query().Get(key)
		if queryValue != "" {
			queryMap[key] = queryValue
		}
	}
	table := persistence.DB.TablesMap["jobs"]
	response, err := table.Search(queryMap)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	} else {
		res, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table := persistence.DB.TablesMap["jobs"]
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

func CreateJob(w http.ResponseWriter, r *http.Request) {
	// create Worker from request body
	var jobMap map[string]string
	json.NewDecoder(r.Body).Decode(&jobMap)
	job, err := models.NewJob(jobMap["device_id"],
		jobMap["reason"],
		jobMap["observations"],
		jobMap["status"],
		jobMap["assigned_id"],
		jobMap["author_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		// add worker to DB
		table := persistence.DB.TablesMap["jobs"]
		response, err := table.Add(job)
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

func UpdateJob(w http.ResponseWriter, r *http.Request) {
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

func DeleteJob(w http.ResponseWriter, r *http.Request) {
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
