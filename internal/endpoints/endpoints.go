package endpoints

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf(w, "Hi there")
}

/**
<<TODO>>
Enpoints to add:
- GetAllDevices() #Retrive all device information from mongodb
- GetDeviceByID() #Retrive device information by id from mongodb
- UpdateDeviceSettings() #Take device settings struct and update the devices settings from (On Device side, will need to send acknowledgment to new setting changes)
- RegisterNewDevice() #Register new device to user. Need to decide how to do registration
- etc.

*/
