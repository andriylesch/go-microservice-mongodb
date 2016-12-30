package models

type User struct {
	Id        string `bson:"_id" json:"id"`
	Nick      string `bson:"nick" json:"nick"`
	Email     string `bson:"email" json:"email"`
	FirstName string `bson:"firstname" json:"firstname"`
	LastName  string `bson:"lastname" json:"lastname"`
	Phone     string `bson:"phone" json:"phone"`
}
