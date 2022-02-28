package greenBot

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"myGreenApi/internal/datastore"
	"myGreenApi/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type GreenDevice interface {
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
	w.Header().Set("Content-Type", "application/json")
	devs, err := gdev.devRepo.FindAll(context.TODO())
	if err != nil {
		log.Printf("Error Occurred getting All Devices: %v", err)
	}

	json.NewEncoder(w).Encode(&devs)
}

func (gdev GreenDevice) DeviceInfo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// we get params with mux.
	var params = mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])
	dev, err := gdev.devRepo.Find(context.TODO(), id)
	if err != nil {
		log.Printf("Error Occurred when reading collection: %v", err)
	}

	json.NewEncoder(w).Encode(&dev)
}

func (gdev GreenDevice) Create(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "application/json" {
		io.WriteString(w, "Content Type is not application/json")
		return
	}
	// RECEIVES: USERID, HUMIDITY, TEMPERATURE.
	params := mux.Vars(req)
	userID, err := primitive.ObjectIDFromHex(params["UserID"])
	if err != nil {
		log.Printf("Error Occurred when creating collection: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	humidity, err := strconv.Atoi(params["humidity"])
	if err != nil {
		log.Printf("Error Occurred when creating collection: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	temp, err := strconv.Atoi(params["temperature"])
	if err != nil {
		log.Printf("Error Occurred when creating collection: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	greenDev := models.Device{
		HumidityLevel: humidity,
		Temperature:   temp,
		UserID:        userID,
	}
	result, err := gdev.devRepo.Create(context.TODO(), &greenDev)
	if err != nil {
		log.Printf("Error Occurred when creating collection: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func (gdev GreenDevice) Delete(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "application/json" {
		io.WriteString(w, "Content Type is not application/json")
	}
	params := mux.Vars(req)
	userID, err := primitive.ObjectIDFromHex(params["UserID"])
	if err != nil {
		log.Printf("Error Occurred when deleting collection: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := gdev.devRepo.Delete(context.TODO(), userID)
	if err != nil {
		log.Printf("Error Occurred when deleting collection: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func (gdev GreenDevice) Update(w http.ResponseWriter, req *http.Request) {
	//
}
