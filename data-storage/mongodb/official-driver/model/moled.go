package model

type Teacher struct {
	FirstName string `bson:"firstname"`
	LastName  string `bson:"lastname"`
}
