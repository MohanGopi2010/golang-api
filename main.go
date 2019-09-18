package main

import (
	"api/apiroutes"
	"api/config"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Load json config file.
	configFile, err := config.LoadapiConfig()

	if err != nil {
		log.Fatal("apiConfig.json file loading error.", err)
		panic(err)

	}

	//Mongo Db Connections
	mConncetion, err := config.MongodbConnect(configFile)

	if err != nil {
		log.Fatal(err)
	} 
	err = mConncetion.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	router := apiroutes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
