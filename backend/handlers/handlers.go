package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func LoadDatabase(databasePath string, tableNames []string) (*persistence.Database, map[string]*persistence.Table) {
	loadDatabase, _ := persistence.LoadDatabase(databasePath)
	tablesMap := map[string]*persistence.Table{}
	for _, tableName := range tableNames {
		table, _ := persistence.LoadTable(databasePath, tableName)
		tablesMap[tableName] = table
	}
	return loadDatabase, tablesMap
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	//databasePath := os.Getenv("DB_PATH")
	databasePath := "./database"
	vars := mux.Vars(r)
	table, _ := persistence.LoadTable(databasePath, "workers")
	response := table.ReadRow(vars["id"])
	responseMap := map[string]string{
		"id":    response[0],
		"name":  response[1],
		"email": response[2],
		"phone": response[3],
	}
	res, _ := json.Marshal(responseMap)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
