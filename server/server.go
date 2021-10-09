package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

type Posts struct {
	USERID string
	ID    string   
	Caption string             
	URL  string           
	Time  string        
}

func users(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	data := User{"1", "ash", "example@gmail.com","admin"}
	b, _ := json.Marshal(data)
	Clientt, _ := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.9w0dh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := Clientt.Database("admin").Collection("admin")
	collection.InsertOne(context.TODO(), data)
	http.Post("localhost:8080/users", "application/json",bytes.NewBuffer(b))
	
}

func posts(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	data := Posts{"1", "2", "golang is hard to learn","googleimages.com","2021-01-01 00:00:01"}
	b, _ := json.Marshal(data)
	Clientt, _ := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.9w0dh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := Clientt.Database("admin").Collection("posts")
	collection.InsertOne(context.TODO(), data)
	http.Post("localhost:8080/posts", "application/json",bytes.NewBuffer(b))
	
}
func Getuser(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	Clientt, _ := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.9w0dh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := Clientt.Database("admin").Collection("admin")
	fmt.Print(collection.FindOne(context.TODO(),request.URL.String()))	

}

func getposts(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	Clientt, _ := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.9w0dh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := Clientt.Database("admin").Collection("posts")
	fmt.Print(collection.FindOne(context.TODO(),request.URL.String()))	

}
func listall(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	Clientt, _ := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.9w0dh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := Clientt.Database("admin").Collection("posts")
	fmt.Print(collection)

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
	http.HandleFunc("/users",users)
	http.HandleFunc("/posts",posts)
	http.HandleFunc("/users/find/",Getuser)
	http.HandleFunc("/posts/find/",getposts)
	http.HandleFunc("/listallposts/",listall)
	http.ListenAndServe(":8080",nil)

}