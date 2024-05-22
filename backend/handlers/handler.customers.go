package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qemanuel/tech-sup-webapp/backend/models"
	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	table := persistence.DB.TablesMap["customers"]
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

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table := persistence.DB.TablesMap["customers"]
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

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// create Customer from request body
	var customerMap map[string]string
	json.NewDecoder(r.Body).Decode(&customerMap)
	customer, err := models.NewCustomer(customerMap["name"], customerMap["email"], customerMap["phone"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		// add customer to DB
		table := persistence.DB.TablesMap["customers"]
		response, err := table.Add(customer)
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

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// create Customer from request body
	var customerMap map[string]string
	json.NewDecoder(r.Body).Decode(&customerMap)
	customer, err := models.NewCustomer(customerMap["name"], customerMap["email"], customerMap["phone"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		// update Customer on DB
		table := persistence.DB.TablesMap["customers"]
		id := vars["id"]
		response, err := table.Update(customer, id)
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

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	table := persistence.DB.TablesMap["customers"]
	vars := mux.Vars(r)
	err := table.Remove(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Customer not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Customer deleted"))
	}
}
