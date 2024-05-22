package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qemanuel/tech-sup-webapp/backend/handlers"
	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func initDatabase(databasePath string) {
	tables := map[string][]string{
		"workers":   {"name", "email", "phone"},
		"customers": {"name", "email", "phone"},
		"devices":   {"owner_id", "kind", "brand", "model", "serial"},
		//		"jobs":       {""},
		//		"incidences": {""},
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
	// workers
	router.HandleFunc("/api/v1/workers/", handlers.GetWorkers).Methods("GET")
	router.HandleFunc("/api/v1/workers/{id:[0-9]+}", handlers.GetWorker).Methods("GET")
	router.HandleFunc("/api/v1/workers/", handlers.CreateWorker).Methods("POST")
	router.HandleFunc("/api/v1/workers/{id:[0-9]+}", handlers.DeleteWorker).Methods("DELETE")
	router.HandleFunc("/api/v1/workers/{id:[0-9]+}", handlers.UpdateWorker).Methods("POST")
	// customers
	router.HandleFunc("/api/v1/customers/", handlers.GetCustomers).Methods("GET")
	router.HandleFunc("/api/v1/customers/{id:[0-9]+}", handlers.GetCustomer).Methods("GET")
	router.HandleFunc("/api/v1/customers/", handlers.CreateCustomer).Methods("POST")
	router.HandleFunc("/api/v1/customers/{id:[0-9]+}", handlers.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/api/v1/customers/{id:[0-9]+}", handlers.UpdateCustomer).Methods("POST")
	// devices
	router.HandleFunc("/api/v1/devices/", handlers.GetDevices).Methods("GET")
	router.HandleFunc("/api/v1/devices/{id:[0-9]+}", handlers.GetDevice).Methods("GET")
	router.HandleFunc("/api/v1/devices/", handlers.CreateDevice).Methods("POST")
	router.HandleFunc("/api/v1/devices/{id:[0-9]+}", handlers.DeleteDevice).Methods("DELETE")
	router.HandleFunc("/api/v1/devices/{id:[0-9]+}", handlers.UpdateDevice).Methods("POST")

	//http.Handle("/", r)
	http.ListenAndServe(":8010", router)
}
