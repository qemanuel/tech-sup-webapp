package handlers

import (
	"encoding/json"
	"fmt"
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
	}
	json.NewEncoder(w).Encode(response)
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
	table := persistence.DB.TablesMap["devices"]
	var device models.Device
	json.NewDecoder(r.Body).Decode(&device)
	id, err := table.Add(device)
	idString := fmt.Sprint(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	device.SetId(idString)
	json.NewEncoder(w).Encode(&device)
}

func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	table := persistence.DB.TablesMap["devices"]
	vars := mux.Vars(r)
	var device models.Device
	json.NewDecoder(r.Body).Decode(&device)
	err := table.Update(device, vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&device)
}

func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	table := persistence.DB.TablesMap["devices"]
	vars := mux.Vars(r)
	err := table.Remove(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Device not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
}
