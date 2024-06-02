package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qemanuel/tech-sup-webapp/backend/handlers"
	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func initDatabase(databasePath string) {
	tables := map[string][]string{
		"workers":    {"name", "email", "phone"},
		"customers":  {"name", "email", "phone"},
		"devices":    {"owner_id", "kind", "brand", "model", "serial"},
		"jobs":       {"device_id", "status", "reason", "observations", "author_id", "assigned_id"},
		"incidences": {"job_id", "body", "next_id", "author_id"},
	}
	persistence.NewDatabase(databasePath)
	for tableName, keys := range tables {
		persistence.DB.NewTable(tableName, keys)
	}
}

func main() {
	//databasePath := os.Getenv("DB_PATH")
	databasePath := "./database"

	initDatabase(databasePath)

	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	// workers
	workers := api.PathPrefix("/workers").Subrouter()
	workers.HandleFunc("/", handlers.GetWorkers).Methods("GET")
	workers.HandleFunc("/{id:[0-9]+}", handlers.GetWorker).Methods("GET")
	workers.HandleFunc("/", handlers.CreateWorker).Methods("POST")
	workers.HandleFunc("/{id:[0-9]+}", handlers.DeleteWorker).Methods("DELETE")
	workers.HandleFunc("/{id:[0-9]+}", handlers.UpdateWorker).Methods("POST")
	// customers
	customers := api.PathPrefix("/customers").Subrouter()
	customers.HandleFunc("/", handlers.GetCustomers).Methods("GET")
	customers.HandleFunc("/{id:[0-9]+}", handlers.GetCustomer).Methods("GET")
	customers.HandleFunc("/", handlers.CreateCustomer).Methods("POST")
	customers.HandleFunc("/{id:[0-9]+}", handlers.DeleteCustomer).Methods("DELETE")
	customers.HandleFunc("/{id:[0-9]+}", handlers.UpdateCustomer).Methods("POST")
	// devices
	devices := api.PathPrefix("/devices").Subrouter()
	devices.HandleFunc("/", handlers.GetDevices).Methods("GET")
	devices.HandleFunc("/{id:[0-9]+}", handlers.GetDevice).Methods("GET")
	devices.HandleFunc("/", handlers.CreateDevice).Methods("POST")
	devices.HandleFunc("/{id:[0-9]+}", handlers.DeleteDevice).Methods("DELETE")
	devices.HandleFunc("/{id:[0-9]+}", handlers.UpdateDevice).Methods("POST")
	// jobs
	jobs := api.PathPrefix("/jobs").Subrouter()
	//Queries("customerId", "{customerId}").
	//Queries("deviceId", "{deviceId}").
	//Queries("authorId", "{authorId}").
	//Queries("assignedId", "{assignedId}").

	jobs.HandleFunc("/", handlers.GetJobs).Methods("GET")
	//Queries("deviceId", "{deviceId}").
	//Queries("authorId", "{authorId}").
	//Queries("assignedId", "{assignedId}")
	//		Queries("customerId", "{customerId}").

	jobs.HandleFunc("/{id:[0-9]+}", handlers.GetJob).Methods("GET")
	jobs.HandleFunc("/", handlers.CreateJob).Methods("POST")
	jobs.HandleFunc("/{id:[0-9]+}", handlers.DeleteJob).Methods("DELETE")
	jobs.HandleFunc("/{id:[0-9]+}", handlers.UpdateJob).Methods("POST")

	//http.Handle("/", r)
	http.ListenAndServe(":8010", api)
}
