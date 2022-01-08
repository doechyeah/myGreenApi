package main

import (
	"context"
	"io"
	"log"
	"myGreenApi/internal/datastore"
	"myGreenApi/internal/endpoints"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var mongoDS *datastore.MongoDataStore

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found. Try again.", http.StatusNotFound)
		return
	}
	err := mongoDS.Session.Ping(context.TODO(), nil)
	var pingOut string
	if err == nil {
		pingOut = "SUCCESS"
	} else {
		pingOut = "FAILED: " + err.Error()
	}
	io.WriteString(w, "Welcome to the GreenAPI!~\nMONGODB Connection: "+pingOut)
}

// func checkDev(w http.ResponseWriter, r *http.Request) {
// 	mongoDS.Session.GetCollection("Devices")
// }

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	URI := os.Getenv("MONGODB_URI")
	if URI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	dbName := os.Getenv("MONGODB_NAME")
	if dbName == "" {
		log.Fatal("You must set your 'MONGODB_NAME' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	config := datastore.NewMongoConfig(URI, dbName)

	mongoDS = datastore.NewDatastore(config)
	defer mongoDS.Session.Disconnect(context.Background())
	// Start HTTP Server
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/", index)
	endpoints.StartDeviceHandlers(apiRouter, mongoDS)
	http.Handle("/", apiRouter)

	http.ListenAndServe(":4210", nil)
}
