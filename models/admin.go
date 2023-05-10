package models


type Admin struct {
	Id 			int
	Username 	string
	Password 	string
}

type Admins []*Admin