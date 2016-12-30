package handlers

import (
	"go-microservice-mongodb/helpers"
	"go-microservice-mongodb/managers"
	"go-microservice-mongodb/models"
	"net/http"
)

const userIdParam = "userId"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := managers.GetUsers()
	if len(users) > 0 {
		helpers.SetResponse(w, http.StatusOK, users)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, nil)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	userId := helpers.GetUserId(r, userIdParam)

	model := managers.GetUserById(userId)

	if (models.User{}) == model {
		helpers.SetResponse(w, http.StatusNotFound, nil)
	} else {
		helpers.SetResponse(w, http.StatusOK, model)
	}
}

// create user
func PostUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	helpers.GetBody(w, r, &user)

	isCreated := managers.InsertUser(user)

	if isCreated {
		helpers.SetResponse(w, http.StatusCreated, nil)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, nil)
	}
}

// update user
func PutUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	userId := helpers.GetUserId(r, userIdParam)
	helpers.GetBody(w, r, &user)

	isCreated := managers.UpdateUser(userId, user)

	if isCreated {
		helpers.SetResponse(w, http.StatusOK, nil)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, nil)
	}
}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	userId := helpers.GetUserId(r, userIdParam)

	isDeleted := managers.DeleteUser(userId)

	if isDeleted {
		helpers.SetResponse(w, http.StatusNoContent, nil)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, nil)
	}
}
