package greenBot

import (
	"context"
	"myGreenApi/internal/datastore"
	"myGreenApi/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const deviceCollection = "Devices"

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

func (r *deviceRepo) Create(ctx context.Context, greenDev *models.Device) (*mongo.InsertOneResult, error) {
	result, err := r.store.DB.Collection(deviceCollection).InsertOne(ctx, &greenDev)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *deviceRepo) Delete(ctx context.Context, userID primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := r.store.DB.Collection(deviceCollection).DeleteOne(ctx, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *deviceRepo) Update(ctx context.Context, greenDev *models.Device) (*mongo.UpdateResult, error) {
	result, err := r.store.DB.Collection(deviceCollection).UpdateByID(ctx, greenDev.ID, greenDev)
	if err != nil {
		return nil, err
	}
	return result, nil
}
