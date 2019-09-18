package Models

import "time"

type GetUsers []GetUser

type GetUser struct {
	Id        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	DOB       time.Time `json:"dob"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	Height    float32   `json:"height"`
	IsWorking bool      `json:"isworking"`
}
