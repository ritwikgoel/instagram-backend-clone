package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
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

//Hashing for the password. Using SHA256 for hashing
func HashIt(b string) string{
	rec:=b
	hh:=sha256.New()
	hh.Write([]byte(rec))
	hashed:=hh.Sum(nil)
	return hex.EncodeToString(hashed)

}

func users(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	data := User{"1", "ash", "example@gmail.com",HashIt("admin")}
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
	//Pushed Password to github. ****** 
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//Connect the client to the DB
	err = Clientt.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//Can omit. At the end, the client will get disconnected. defer acts like a stack.LIFO
	defer Clientt.Disconnect(ctx)
	//Pinging the client. Making sure about connections
	err = Clientt.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	//Handing routes
	http.HandleFunc("/users",users)
	http.HandleFunc("/posts",posts)
	http.HandleFunc("/users/:find/",Getuser)
	http.HandleFunc("/posts/:find/",getposts)
	http.HandleFunc("/listallposts/",listall)
	http.ListenAndServe(":8080",nil)//Listening on port 8080
}