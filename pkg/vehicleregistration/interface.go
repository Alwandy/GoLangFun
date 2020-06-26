//Package vregistration interface
package vregistration

//VehicleRegistrations struct
type VehicleRegistrations struct {
	Registrations []VehicleRegistration `json:"registrations"`
}

//VehicleRegistration struct
type VehicleRegistration struct {
	Registration string `json:"registration"`
	Brand        string `json:"brand"`
	Colour       string `json:"colour"`
	Owner        string `json:"owner"`
}

//Path to json
const filePath = "data/vehicleregistration.json"
