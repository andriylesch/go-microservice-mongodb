# go-microservice-mongodb

How to create microservice and make friends GO and MongoDb

##Install

    go get https://github.com/andriylesch/go-microservice-mongodb

##Install MongoDb

For this application you can install MongoDb on your local system or use docker. 
In current application MongoDb was installed on local system. 
How you can do it, you can find detail information 
https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/

##Building a REST API

For full CRUD logic was implemented list of methods :

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

Endpoints for GetUsers method :
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

Endpoints for GetUserById method :
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

Endpoints for PostUser method :
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

Endpoints for PostUser method :
```go
    http://localhost:8080/users -POST
```

result will be as
```go
  Http status : 201 (Created)
```






- 








