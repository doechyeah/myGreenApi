package greenBot

import (
	"context"
	"myGreenApi/internal/datastore"
	"myGreenApi/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var deviceCollection string = "Devices"

type DeviceRepo interface {
	Find(ctx context.Context, id primitive.ObjectID) (*models.Device, error)
}

type deviceRepo struct {
	store *datastore.MongoDataStore
}

func (r *deviceRepo) Find(ctx context.Context, id primitive.ObjectID) (*models.Device, error) {
	var greenDev models.Device
	err := r.store.DB.Collection(deviceCollection).FindOne(ctx, bson.M{"_id": id}).Decode(&greenDev)
	if err != nil {
		return nil, err
	}
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &greenDev, nil
}
