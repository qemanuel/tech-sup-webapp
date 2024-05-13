package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qemanuel/tech-sup-webapp/backend/handlers"
	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func initDatabase(databasePath string) {
	tables := map[string][]string{
		"workers":   {"id", "name", "email", "phone"},
		"customers": {"id", "name", "email", "phone"},
		"devices":   {"id", "owner", "kind", "brand", "model", "serial"},
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
	router.HandleFunc("/api/v1/workers/", handlers.GetWorkers).Methods("GET")
	router.HandleFunc("/api/v1/workers/{id:[0-9]+}", handlers.GetWorker).Methods("GET")
	router.HandleFunc("/api/v1/workers/", handlers.CreateWorker).Methods("POST")
	router.HandleFunc("/api/v1/workers/{id:[0-9]+}", handlers.DeleteWorker).Methods("DELETE")
	router.HandleFunc("/api/v1/workers/{id:[0-9]+}", handlers.UpdateWorker).Methods("POST")

	//http.Handle("/", r)
	http.ListenAndServe(":8010", router)
}
