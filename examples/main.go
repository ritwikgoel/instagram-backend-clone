package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
type User struct {
	ID    string   
	Name string             
	Email  string           
	Password  string        
}

func users(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	data := User{"1", "ash", "example@gmail.com","admin"}
	b, _ := json.Marshal(data)


	http.Post("localhost:8080/users", "application/json",bytes.NewBuffer(b))
	
}
func getuser(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	Clientt, _ := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.9w0dh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := Clientt.Database("admin").Collection("admin")
	fmt.Print(collection.FindOne(context.TODO(),request.URL.String()))	

}

func main() {
	fmt.Print("init")
	Clientt, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.9w0dh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	//Yes i know i pushed the password. 
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Clientt.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer Clientt.Disconnect(ctx)
	err = Clientt.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)
	
	// data := User{"1", "ash", "example@gmail.com","admin"}
	// collection.InsertOne(context.TODO(), data)
	// if err != nil {
    // 	log.Fatal(err)
	// }
	http.HandleFunc("/users",users)
	http.HandleFunc("/users/find/",getuser)
	http.ListenAndServe(":8080",nil)

}