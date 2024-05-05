package main

import (
	"fmt"

	"github.com/qemanuel/tech-sup-webapp/backend/models"
	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func main() {

	worker1, _ := models.NewWorker("worker1",
		"worker1@gmail.com",
		"152345424")

	worker2, _ := models.NewWorker("worker2",
		"worker2@gmail.com",
		"124453453")

	worker3, _ := models.NewWorker("worker3",
		"worker3@gmail.com",
		"124453453")

	workersTable, _ := persistence.NewTable("workers", []string{"id", "name", "email", "phone"})

	workersTable.AddElement(worker1.StringWorker())
	workersTable.AddElement(worker2.StringWorker())
	workersTable.AddElement(worker3.StringWorker())

	fmt.Println(workersTable.GetElement(2))

	workersTable.RemoveElement(1)
	fmt.Println(workersTable.GetAllElements())

}
