# go-microservice-mongodb

How to create microservice and make friends GO and MongoDb

##Install

    go get https://github.com/andriylesch/go-microservice-mongodb

##Install MongoDb

For this application you can install MongoDb on your local system or use docker. 
In current application MongoDb was installed on local system. 
How you can do it, you can find detail information 
https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/

##Generate JSON 
There is [JSON generator] (http://www.json-generator.com/) website. which allows to developer generate random JSON. 
For this test app was created JSON template :
```go
[
  '{{repeat(5, 7)}}',
  {
    _id: '{{objectId()}}',
    nick : '{{company()}}',
    email: '{{email()}}',
    firstname: '{{firstName()}}',
    lastname: '{{surname()}}',
    phone: '+1 {{phone()}}'
  }
]
```

##Building a REST API

All CRUD logic use one custom `User` type:

```go
type User struct {
	Id        string `bson:"_id" json:"id"`
	Nick      string `bson:"nick" json:"nick"`
	Email     string `bson:"email" json:"email"`
	FirstName string `bson:"firstname" json:"firstname"`
	LastName  string `bson:"lastname" json:"lastname"`
	Phone     string `bson:"phone" json:"phone"`
}
```

List of methods :

```go

const userIdParam = "userId"

func GetUsers(w http.ResponseWriter, r *http.Request) { ... }

func GetUserById(w http.ResponseWriter, r *http.Request)  { ... }

// create user
func PostUser(w http.ResponseWriter, r *http.Request)  { ... }

// update user
func PutUser(w http.ResponseWriter, r *http.Request)  { ... }

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request)  { ... }

```

Endpoints for `GetUsers` method :
```go
    http://localhost:8080/users -GET
```

result will be as
```json
[
    {
        "id": "586555b6fcd853a100e0c77e",
        "nick": "Zyple",
        "email": "bishopdillon@zyple.com",
        "firstname": "Petra",
        "lastname": "Mathis",
        "phone": "+1 (893) 574-2359"
    },
    ...
]    
```
#

Endpoints for `GetUserById` method :
```go
    http://localhost:8080/users/<userId> -GET
```

result will be as
```json
    {
        "id": "586555b6fcd853a100e0c77e",
        "nick": "Zyple",
        "email": "bishopdillon@zyple.com",
        "firstname": "Petra",
        "lastname": "Mathis",
        "phone": "+1 (893) 574-2359"
    } 
```
#

Endpoints for `PostUser` (Insert new user) method :
```go
    http://localhost:8080/users -POST
```
JSON parameter
```json
{
  "nick": "testnick",
  "email": "test@gmail.com",
  "firstname": "testFN",
  "lastname": "testLN",
  "phone": "+0 (000) 123-4567"
}
```

result will be as
```go
  Http status : 201 (Created)
```
#

Endpoints for `PutUser` (Update user) method :
```go
    http://localhost:8080/users/<userId> -PUT
```

JSON parameter
```go
{
   "nick": "testnick123",
   "email": "test@gmail.com",
   "firstname": "testFN",
   "lastname": "testLN",
   "phone": "+0 (000) 000-123445"
}
```

result will be as
```go
  Http status : 200 (OK)
```

#

Endpoints for `DeleteUser` method :
```go
    http://localhost:8080/users/<userId> -DELETE
```

result will be as
```go
  Http status : 204 (no content)
```

##Reference

JSON generator
- http://www.json-generator.com/

MongoDB
- https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/
- https://mongobooster.com/downloads - allows to manage MongoDB. 

Go language
- http://thenewstack.io/make-a-restful-json-api-go/
- https://stevenwhite.com/building-a-rest-service-with-golang-1/
- https://stevenwhite.com/building-a-rest-service-with-golang-2/
- https://stevenwhite.com/building-a-rest-service-with-golang-1/

