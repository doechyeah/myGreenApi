package endpoints

import (
	"myGreenApi/internal/datastore"
	"myGreenApi/internal/endpoints/greenBot"
	"myGreenApi/internal/endpoints/user"

	"github.com/gorilla/mux"
)

/**
<<TODO>>
Enpoints to add:
- GetAllDevices() #Retrive all device information from mongodb
- GetDeviceByID() #Retrive device information by id from mongodb
- UpdateDeviceSettings() #Take device settings struct and update the devices settings from (On Device side, will need to send acknowledgment to new setting changes)
- RegisterNewDevice() #Register new device to user. Need to decide how to do registration
- etc.
*/

func StartDeviceHandlers(router *mux.Router, mongoDS *datastore.MongoDataStore) {
	greenDev := greenBot.NewGreenDevice(mongoDS)
	router.HandleFunc("/device", greenDev.Devices).Methods("GET")
	router.HandleFunc("/device", greenDev.Create).Methods("PUT")
	router.HandleFunc("/device/{id}", greenDev.DeviceInfo).Methods("GET")
	router.HandleFunc("/device/{id}", greenDev.Update).Methods("POST")
	router.HandleFunc("/device/{id}", greenDev.Delete).Methods("DELETE")
}

func StartUserHandlers(router *mux.Router) {
	router.HandleFunc("/user", user.UserInfo).Methods("GET")
	router.HandleFunc("/user", user.Create).Methods("PUT")
	router.HandleFunc("/user", user.Update).Methods("POST")
	router.HandleFunc("/user/{deviceID}", user.Delete).Methods("DELETE")
}
