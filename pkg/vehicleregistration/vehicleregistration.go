//Package vregistration class
package vregistration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ReadDatabase returns the json data unmarshalled into the struct vehicleRegistration{}
// No values required to populate
func (v *VehicleRegistrations) ReadDatabase() VehicleRegistrations {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database found")
	defer file.Close() // Always close after loading!

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	var vehicleregistrations VehicleRegistrations

	// unmarshall it
	err = json.Unmarshal(byteValue, &vehicleregistrations)
	if err != nil {
		fmt.Println("error:", err)
	}

	return vehicleregistrations
}

// AddNewVehicle will append to existing struct from vehicleregistrations.ReadDatabase
// Before appending it will loop through to see if value exists. (Can we improve this? What is the time and space complexity if there's over 100 entries?)
// Returns string if successful or not (Could return bool / error)
func AddNewVehicle(registration, brand, colour, owner string) error {
	var database VehicleRegistrations  // Define database to use struct
	database = database.ReadDatabase() // Populate the struct with the values

	// Lets create the struct to append into existing database
	data := VehicleRegistration{
		Registration: registration,
		Brand:        brand,
		Colour:       colour,
		Owner:        owner,
	}

	for i := range database.Registrations {
		if database.Registrations[i].Registration == data.Registration {
			return fmt.Errorf("Registration already exists, please update instead of adding")
		}
	}

	// Should we check for empty strings or invalid characters?

	database.Registrations = append(database.Registrations, data)
	file, err := json.MarshalIndent(database, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile(filePath, file, 0644)
	fmt.Println("Registration has been added into system.")
	return nil
}
