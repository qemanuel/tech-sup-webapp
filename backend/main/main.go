package main

import (
	"fmt"

	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func main() {

	//	worker1, _ := models.NewWorker("worker1",
	//		"worker1@gmail.com",
	//		"152345424")
	//
	//	worker2, _ := models.NewWorker("worker2",
	//		"worker2@gmail.com",
	//		"124453453")
	//
	//	worker3, _ := models.NewWorker("worker3",
	//		"worker3@gmail.com",
	//		"124453453")
	//
	databasePath := "/Users/qemanuel/IT/personal/github/tech-sup-webapp/backend/main/database"
	//database, _ := persistence.NewDatabase(databasePath)
	//workersTable, _ := database.NewTable("workers", []string{"id", "name", "email", "phone"})

	database, _ := persistence.LoadDatabase(databasePath)
	workersTable, _ := database.NewTable("test", []string{"id", "name", "email", "phone"})
	fmt.Println(workersTable)

	//	workersTable.AddElement(worker1.StringWorker())
	//	workersTable.AddElement(worker2.StringWorker())
	//	workersTable.AddElement(worker3.StringWorker())

	//fmt.Println(workersTable.GetElement(2))
	//workersTable.RemoveElement(1)
	//fmt.Println(workersTable.GetAllElements())

}
