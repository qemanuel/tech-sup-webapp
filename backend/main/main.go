package main

import (
	"fmt"
	"net/http"
	"os"

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
	tablesMap := map[string]*persistence.Table{}
	newDatabase, _ := persistence.NewDatabase(databasePath)
	for tableName, keys := range tables {
		table, _ := newDatabase.NewTable(tableName, keys)
		tablesMap[tableName] = table
	}
}

func main() {
	//databasePath := os.Getenv("DB_PATH")
	databasePath := "./database"
	_, err := os.Stat(fmt.Sprintf("%s/database.csv", databasePath))
	if err != nil {
		initDatabase(databasePath)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/workers/{id:[0-9]+}", handlers.GetCustomer)
	//http.Handle("/", r)
	http.ListenAndServe(":8010", router)
}
