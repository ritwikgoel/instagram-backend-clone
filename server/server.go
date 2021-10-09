package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
	Email  string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
}
var client *mongo.Client


func homePage(response http.ResponseWriter, request *http.Request){
	fmt.Fprint(response,"Home Page")
}

func createUser(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("ritwikgoel").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
	
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user User
	collection := client.Database("ritwikgoel").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection.Find(ctx, bson.M{})
	json.NewEncoder(response).Encode(user)
}


func main(){
	fmt.Print("Starting")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	http.HandleFunc("/",homePage)
	//createUserr:=createUser()
	http.HandleFunc("/createUser",createUser)
	http.HandleFunc("/GetUser",GetUser)
	http.ListenAndServe(":8080",nil)
}