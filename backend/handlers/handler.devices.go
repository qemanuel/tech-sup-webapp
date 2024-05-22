package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qemanuel/tech-sup-webapp/backend/models"
	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func GetDevices(w http.ResponseWriter, r *http.Request) {
	table := persistence.DB.TablesMap["devices"]
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

func GetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table := persistence.DB.TablesMap["devices"]
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

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	// create Device from request body
	var deviceMap map[string]string
	json.NewDecoder(r.Body).Decode(&deviceMap)
	device, err := models.NewDevice(deviceMap["brand"],
		deviceMap["kind"],
		deviceMap["model"],
		deviceMap["owner_id"],
		deviceMap["serial"],
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		// add Device to DB
		table := persistence.DB.TablesMap["devices"]
		response, err := table.Add(device)
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

func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// create Device from request body
	var deviceMap map[string]string
	json.NewDecoder(r.Body).Decode(&deviceMap)
	device, err := models.NewDevice(deviceMap["brand"],
		deviceMap["kind"],
		deviceMap["model"],
		deviceMap["owner_id"],
		deviceMap["serial"],
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		// update Device on DB
		table := persistence.DB.TablesMap["devices"]
		id := vars["id"]
		response, err := table.Update(device, id)
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

func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	table := persistence.DB.TablesMap["devices"]
	vars := mux.Vars(r)
	err := table.Remove(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Device not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Device deleted"))
	}
}
