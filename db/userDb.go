package db

import (
	"api/Models"
	"api/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Init : Initilizing the log	
func Init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

//GetAllUsers :
func GetAllUsers() (_allusers []Models.GetUser, err error) {

	//Check whether mongo db is able to connect or not
	err = config.MongoDBConnection.Ping(context.TODO(), nil)
	if err != nil {
		log.Panicln("Data source is not connected.")
		panic(err)
	}

	fmt.Println(config.ApiConfigAL.DB)

	//get the data from DB
	//TODO: Need to check whether DB is exists or not.
	_userCollections := config.MongoDBConnection.Database(config.ApiConfigAL.DB).Collection("user")
	filterOptions := options.Find()

	curr, err := _userCollections.Find(context.TODO(), filterOptions)
	if err != nil {
		log.Fatalf("Error while gettting the user data. %v", err)
	}

	for curr.Next(context.TODO()) {
		var _eachUser Models.GetUser

		err := curr.Decode(&_eachUser)

		if err != nil {
			log.Fatalf("Error while reading each data. %v", err)
		}

		_allusers = append(_allusers, _eachUser)

	}

	curr.Close(context.TODO())

	return _allusers, err

}

//GetUser : returning specific user based on ID match
func GetUser(id int) (_user Models.GetUser, err error) {

	var _userData Models.GetUser
	_collections := config.MongoDBConnection.Database(config.ApiConfigAL.DB).Collection("user")
	_filterOptions := bson.M{"id": id}
	_collections.FindOne(context.TODO(), _filterOptions).Decode(&_userData)
	return _userData, nil
}

// Createuser ... Create new user
func Createuser(_createUser Models.GetUser) (_id int, err error) {

	_collections := config.MongoDBConnection.Database(config.ApiConfigAL.DB).Collection("user")
	log.Println(_collections)
	_, err = _collections.InsertOne(context.TODO(), _createUser)

	if err != nil {
		log.Fatal(err)
	}
	return _createUser.Id, err

}

//UpdateUser ...
func UpdateUser(_id int, _updateUser Models.GetUser) (_returnID int, err error) {

	_collections := config.MongoDBConnection.Database(config.ApiConfigAL.DB).Collection("user")
	_filterOptions := bson.M{"id": _id}
	updateQry := bson.M{"$set": bson.M{
		"firstname": _updateUser.Firstname,
		"lastname":  _updateUser.Lastname,
		"dob":       _updateUser.DOB,
		"height":    _updateUser.Height,
		"isworking": _updateUser.IsWorking,
		"username":  _updateUser.UserName,
		"password":  _updateUser.Password,
	}}
	_, err = _collections.UpdateOne(context.TODO(), _filterOptions, updateQry, nil)

	if err != nil {
		log.Fatal(err)
	}

	return _id, err

}

//DeleteUser ...
func DeleteUser(_id int) (_returnID int, err error) {
	_collections := config.MongoDBConnection.Database(config.ApiConfigAL.DB).Collection("user")
	_filterOptions := bson.M{"id": _id}
	_, err = _collections.DeleteOne(context.TODO(), _filterOptions)

	if err != nil {
		log.Fatal(err)
	}
	return _id, err
}
