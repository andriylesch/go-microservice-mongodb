package managers

import (
	"go-microservice-mongodb/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetUsers() []models.User {

	var users []models.User

	// build query for get all users
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&users)
	}

	executeQuery(userCollectionName, query)
	return users
}

func GetUserById(userId string) models.User {

	var user models.User

	if !bson.IsObjectIdHex(userId) {
		return user
	}

	query := func(c *mgo.Collection) error {
		return c.FindId(userId).One(&user)
	}

	executeQuery(userCollectionName, query)

	return user
}

func InsertUser(user models.User) bool {

	result := true

	user.Id = bson.NewObjectId().Hex()
	query := func(c *mgo.Collection) error {
		return c.Insert(user)
	}

	err := executeQuery(userCollectionName, query)
	if err != nil {
		result = false
	}

	return result
}

func UpdateUser(userId string, user models.User) bool {

	result := true

	user.Id = userId

	query := func(c *mgo.Collection) error {
		return c.UpdateId(userId, user)
	}

	err := executeQuery(userCollectionName, query)
	if err != nil {
		result = false
	}

	return result
}

func DeleteUser(userId string) bool {
	result := true

	query := func(c *mgo.Collection) error {
		return c.RemoveId(userId)
	}

	err := executeQuery(userCollectionName, query)
	if err != nil {
		result = false
	}

	return result
}
