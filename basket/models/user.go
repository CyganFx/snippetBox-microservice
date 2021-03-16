package models

type User struct {
	ID           int    `json:"id"`
	Firstname    string `json:"firstname"`
	Mobileno     int    `json:"mobile_number"`
	Emailid      string `json:"email"`
	Password     string `json:"password"`
	Createddate  string
	Modifieddate string
	Updateddate  string
}
