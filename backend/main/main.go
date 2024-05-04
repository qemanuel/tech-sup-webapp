package main

import (
	"fmt"

	"github.com/qemanuel/tech-sup-webapp/backend/models"
)

func main() {
	worker, _ := models.NewWorker("ema",
		"ema@gmail.com",
		"152345424")
	customer, _ := models.NewCustomer("google",
		"google@google.com",
		"1234567435")
	device, _ := models.NewDevice(1,
		customer,
		"joystick",
		"sony",
		"dualshock")
	list := models.NewIncidenceList()
	list.AddIncidence(device, "ingressed device", worker)
	inc, _ := list.AddIncidence(device, "cleaning device", worker)
	list.AddIncidence(device, "testing device", worker)
	fmt.Println(list.GetAllIncidences())
	list.RemoveIncidence(inc)
	fmt.Println(list.GetAllIncidences())
}
