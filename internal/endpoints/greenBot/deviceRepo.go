package greenBot

import (
	"context"
	"errors"
	"myGreenApi/internal/datastore"
	"myGreenApi/internal/models"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var deviceCollection string = "Devices"

// type DeviceRepo interface {
// 	Find(ctx context.Context, id primitive.ObjectID) (*models.Device, error)
// }

type deviceRepo struct {
	store *datastore.MongoDataStore
}

func (r *deviceRepo) FindAll(ctx context.Context) ([]models.Device, error) {
	var greenDevs []models.Device
	curr, err := r.store.DB.Collection(deviceCollection).Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for curr.Next(ctx) {
		var greenDev models.Device
		err := curr.Decode(&greenDev)
		if err != nil {
			return nil, err
		}
		greenDevs = append(greenDevs, greenDev)
	}
	return greenDevs, nil
}

func (r *deviceRepo) Find(ctx context.Context, id primitive.ObjectID) (*models.Device, error) {
	var greenDev models.Device
	err := r.store.DB.Collection(deviceCollection).FindOne(ctx, bson.M{"_id": id}).Decode(&greenDev)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &greenDev, nil
}

func (r *deviceRepo) Create(ctx context.Context, params map[string]string) (*mongo.InsertOneResult, error) {
	if params["UserID"] == "" {
		return nil, errors.New("No UserID Provided")
	}
	userID, err := primitive.ObjectIDFromHex(params["UserID"])
	if err != nil {
		return nil, err
	}
	humidity, err := strconv.Atoi(params["humidity"])
	if err != nil {
		return nil, err
	}
	temp, err := strconv.Atoi(params["temperature"])
	if err != nil {
		return nil, err
	}
	greenDev := models.Device{
		HumidityLevel: humidity,
		Temperature:   temp,
		UserID:        userID,
	}
	result, err := r.store.DB.Collection(deviceCollection).InsertOne(ctx, &greenDev)
	if err != nil {
		return nil, err
	}
	return result, nil
}
