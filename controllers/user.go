package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Debsnil24/GO_Mongo.git/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	session *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid User Input")
		return
	}
	filter := bson.M{"_id": oid}
	u := models.User{}
	fmt.Println(filter)
	if err := uc.session.Database("mongo_go").Collection("users").FindOne(context.Background(), filter).Decode(&u); err != nil {
		if err == mongo.ErrNoDocuments {
            w.WriteHeader(http.StatusNotFound) // 404 Not Found
            fmt.Fprintln(w, "User not found")
        } else {
            w.WriteHeader(http.StatusInternalServerError) // 500 for other errors
            fmt.Fprintln(w, "Error querying the database")
            fmt.Println("Database error:", err)
        }
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.ID = primitive.NewObjectID()
	uc.session.Database("mongo_go").Collection("users").InsertOne(context.TODO(), u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid User ID")
		return
	}

	dr, err := uc.session.Database("mongo_go").Collection("users").DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error deleting user")
		return
	}
	if dr.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found")
		return
	}
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted User: %s\n", oid.Hex())
}