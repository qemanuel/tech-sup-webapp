package handlers

import (
	"encoding/json"
	"fmt"
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
	}
	json.NewEncoder(w).Encode(response)
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
	table := persistence.DB.TablesMap["customers"]
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	id, err := table.Add(customer)
	idString := fmt.Sprint(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	customer.SetId(idString)
	json.NewEncoder(w).Encode(&customer)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	table := persistence.DB.TablesMap["customers"]
	vars := mux.Vars(r)
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	err := table.Update(customer, vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&customer)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	table := persistence.DB.TablesMap["customers"]
	vars := mux.Vars(r)
	err := table.Remove(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Customer not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
}
