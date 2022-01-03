package greenBot

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"myGreenApi/internal/datastore"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type GreenDeviceInterface interface {
// 	Devices()
// 	DeviceInfo()
// 	Create()
// 	Delete()
// 	Update()
// }

type GreenDevice struct {
	devRepo deviceRepo
}

func NewGreenDevice(store *datastore.MongoDataStore) GreenDevice {
	devRepo := deviceRepo{store: store}
	return GreenDevice{devRepo: devRepo}
}

func (gdev GreenDevice) Devices(w http.ResponseWriter, req *http.Request) {
	// TEMP
	log.Print("RETURN ALL DEVICES")
	io.WriteString(w, "RETURN ALL DEVICES")
}

func (gdev GreenDevice) DeviceInfo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// we get params with mux.
	var params = mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])
	log.Printf("Device ID Being Asked: %v", id)
	dev, err := gdev.devRepo.Find(context.TODO(), id)
	if err != nil {
		log.Fatalf("Error Occurred when reading collection: %v", err)
	}

	json.NewEncoder(w).Encode(&dev)
}

func (gdev GreenDevice) Create(w http.ResponseWriter, req *http.Request) {
	// TEMP
	io.WriteString(w, "CREATE NEW DEVICE")
}

func (gdev GreenDevice) Delete(w http.ResponseWriter, req *http.Request) {
	// TEMP
}

func (gdev GreenDevice) Update(w http.ResponseWriter, req *http.Request) {
	//temp
}
