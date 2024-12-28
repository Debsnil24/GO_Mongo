package main

import (
    "context"
    "log"
    "net/http"

    "github.com/Debsnil24/GO_Mongo.git/controllers"
    "github.com/julienschmidt/httprouter"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    r := httprouter.New()
    uc := controllers.NewUserController(getSession())
    r.GET("/user/:id", uc.GetUser)
    r.POST("/user", uc.CreateUser)
    r.DELETE("/user/:id", uc.DeleteUser)
    log.Fatal(http.ListenAndServe("localhost:9000", r))
}

func getSession() *mongo.Client {
    ctx := context.Background()
	options := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(ctx, options)
    if (err != nil) {
        panic(err)
    }
    err = client.Ping(ctx, nil)
    if (err != nil) {
        panic(err)
    }
    return client
}