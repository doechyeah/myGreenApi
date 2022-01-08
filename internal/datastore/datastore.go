package datastore

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI          string
	DatabaseName string
}

func NewMongoConfig(uri_cred string, dbName string) *MongoConfig {
	config := MongoConfig{URI: uri_cred, DatabaseName: dbName}
	return &config
}

type MongoDataStore struct {
	DB      *mongo.Database
	Session *mongo.Client
}

func NewDatastore(config *MongoConfig) *MongoDataStore {
	var mongoDataStore *MongoDataStore
	db, session := connect(config)
	if db != nil && session != nil {
		mongoDataStore = new(MongoDataStore)
		mongoDataStore.DB = db
		mongoDataStore.Session = session
		return mongoDataStore
	}
	log.Fatal("Failed to Connect to Database")
	return nil
}

func connect(config *MongoConfig) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo(config)
	})

	return db, session
}

func connectToMongo(config *MongoConfig) (a *mongo.Database, b *mongo.Client) {

	var err error
	session, err := mongo.NewClient(options.Client().ApplyURI(config.URI))
	if err != nil {
		log.Fatal(err)
	}
	session.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	var DB = session.Database(config.DatabaseName)
	log.Printf("CONNECTED TO %v", config.DatabaseName)

	return DB, session
}
