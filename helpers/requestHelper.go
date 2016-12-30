package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"go-microservice-mongodb/models"

	"github.com/gorilla/mux"
)

func GetUserId(request *http.Request, paramName string) string {

	vars := mux.Vars(request)
	var userID string
	userID = vars[paramName]

	return userID
}

func GetBody(response http.ResponseWriter, request *http.Request, user *models.User) {

	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(response).Encode(err); err != nil {
			panic(err)
		}
	}

}
