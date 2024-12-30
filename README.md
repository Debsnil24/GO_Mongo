# GO_Mongo

This is a Golang project that uses a MongoDB backend server & REST API to manage User Information


## Roadmap

### Project Structure
![ProjectStructure](https://github.com/user-attachments/assets/a74ae613-7ee8-4d35-b722-a81cb08ae4ae)

### Routes
![Routes](https://github.com/user-attachments/assets/7809eb27-8907-4f41-bcd5-23f7842f8c0c)


## Testing Demo

Refer to the provided Google Doc for the [Testing Demo](https://docs.google.com/document/d/e/2PACX-1vRXnNwirXxLP5v6zcG6VYxRFBDxoofnipU67XdWlSBuXQCtzAmkME3HO-GmqbTAPZV1nSYkCReyP1CI/pub)




## Run Locally

Clone the project

```bash
  git clone https://github.com/Debsnil24/GO_Mongo.git
```

Go to the project directory

```bash
  cd {Directory}/GO_Mongo
```

Install dependencies

```go
  go get "github.com/julienschmidt/httprouter"
```
```go
  go get "go.mongodb.org/mongo-driver/mongo"
```
```go
  go get "go.mongodb.org/mongo-driver/mongo/options"
```
```go
  go get "go.mongodb.org/mongo-driver/bson"
```
```go
  go get "go.mongodb.org/mongo-driver/bson/primtive"
```
Starting the server

- Connect to the local instance of MongoDB Server
- Add the URI of the server to getSession function in the main.go file 

- Build the project
```go
  go build main.go
```
- Run the main.go file
```go
  go run main.go
```

