package models

import "time"

//Credentials capture user credentials
type Credentials struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordHash string
}

//Masterplan stores retrived from DB to be written in csv file
type Masterplan struct {
	SlNo      string
	Activity  string
	StartDate time.Time
	EndDate   time.Time
}
