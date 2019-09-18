package handler

import (
	"api/Models"
	"api/db"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllUsers :
func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_allUsersDB, err := db.GetAllUsers()
	if err != nil {
		log.Fatalln("Error: db.GetAllUsers() get data from datasource.")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: error occured while reading the data from DB."))
	}

	if _allUsersDB != nil {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(_allUsersDB); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: error occured while reading the data from DB."))
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

}

//GetByID :
func GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_reqValues := mux.Vars(r)
	_id, err := strconv.Atoi(_reqValues["Id"])

	if err != nil {
		log.Fatal("GetUser id should not be empty. ")
		http.Error(w, "Input: Id is mandatory.", http.StatusBadRequest)
		return
	}

	if _id > 0 {

		_getUser, err := db.GetUser(_id)

		if err != nil {
			log.Fatalln("Error: db.GetUser() get data from datasource.")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops something went wrong in data source level."))
			return
		} else {

			if _getUser == (Models.GetUser{}) {
				w.WriteHeader(http.StatusNoContent)
				return
			} else {
				w.WriteHeader(http.StatusOK)
				if err = json.NewEncoder(w).Encode(_getUser); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Error: Error occured while deserilizing the data."))

				}
				return
			}
		}

	} else {
		w.WriteHeader(http.StatusNoContent)
	}

}

//CreateUser :
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Input: Body content should not be empty."))
		return
	}

	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: While reading body data error occured."))
		return
	}

	var _getUser Models.GetUser

	if err := json.Unmarshal(_body, &_getUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: While deseriliing the data error occured."))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if (Models.GetUser{}) == _getUser {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Input: Body content data is mandatory."))
		return
	}

	//log.Print(_getUser)

	_newID, err := db.Createuser(_getUser)

	if err != nil {
		log.Fatalln("Error: db.CreateUser() error occured in create user.")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write([]byte("Oops something went wrong while creating create user data source."))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)

		if err = json.NewEncoder(w).Encode(_newID); err != nil {
			log.Fatalln("Error: handler.CreateUser() error occured while deserilizing the data.")
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("Oops something went wrong while creating create user data source."))
			return
		}
		return
	}

}

//UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_reqValues := mux.Vars(r)
	_id, err := strconv.Atoi(_reqValues["Id"])
	if err != nil {
		log.Fatal("Id should not be empty. ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal("In update body content should not be empty. ", err)
		http.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}

	if err := r.Body.Close(); err != nil {
		http.Error(w, "Error while deserilizing the input data.", http.StatusInternalServerError)
		return
	}

	var _getUser Models.GetUser

	if err := json.Unmarshal(_body, &_getUser); err != nil {
		http.Error(w, "Error while deserilizing the input data.", http.StatusInternalServerError)
		return
	}

	_, err = db.UpdateUser(_id, _getUser)
	if err != nil {
		log.Fatalln("Error: db.CreateUser() error occured in create user.")
		http.Error(w, "Error: db.CreateUser() error occured in create user.", http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
		return
	}

}

//DeleteUser:
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_reqValues := mux.Vars(r)
	_id, err := strconv.Atoi(_reqValues["Id"])

	if err != nil {
		log.Fatal("Id should not be empty. ", err)
		http.Error(w, "Id is mandatory.", http.StatusBadRequest)
		return
	}

	_, err = db.DeleteUser(_id)
	if err != nil {
		log.Fatalln("Error: db.CreateUser() error occured in create user.")
		http.Error(w, "Error: db.Deleteuser() error occured in delete user.", http.StatusBadRequest)
		return

	} else {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
