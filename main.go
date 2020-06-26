package main

import (
	"fmt"

	VRegistration "./pkg/vehicleregistration"
)

func main() {
	//Add a new vehicle
	err := VRegistration.AddNewVehicle("", "", "", "")
	if err != nil {
		fmt.Printf("[ERROR]:%v", err)
	}
}
