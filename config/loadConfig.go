package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBConnection *mongo.Client

var ApiConfigAL ApiConfig

// type MongodbConfig struct {
// 	Server string `json:"server"`
// 	Port   int    `json:"port"`
// }

type ApiConfig struct {
	//mongodbConfig MongodbConfig `json:"mongodbconfig"`
	Server string `json:"server"`
	Port   int    `json:"port"`
	DB     string `json:"db"`
}

func LoadapiConfig() (_config ApiConfig, err error) {
	var _readConfig ApiConfig
	//Read the config json file
	configFile, err := os.Open("D://SampleProjects/GO/src/api/apiConfig.json")
	defer configFile.Close()

	if err != nil {
		return _readConfig, err
	}
	_jsonByteValue, _ := ioutil.ReadAll(configFile)

	json.Unmarshal([]byte(_jsonByteValue), &_readConfig)

	fmt.Println(_readConfig)

	ApiConfigAL = _readConfig

	return _readConfig, err
}

func MongodbConnect(_config ApiConfig) (_connection *mongo.Client, err error) {
	if MongoDBConnection == nil {

		mongoDbConnectionOptions := options.Client().ApplyURI("mongodb://" + _config.Server + ":" + strconv.Itoa(_config.Port))
		mConncetion, err := mongo.Connect(context.TODO(), mongoDbConnectionOptions)

		if err != nil {
			log.Fatal(err)
		}

		MongoDBConnection = mConncetion

		err = mConncetion.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal("Error while connecting mongodb. ", err)
			//panic(err)
		}

	}

	return MongoDBConnection, err

}
