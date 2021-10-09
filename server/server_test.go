package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func getuser(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	Clientt, _ := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.9w0dh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := Clientt.Database("admin").Collection("admin")
	fmt.Print(collection.FindOne(context.TODO(),request.URL.String()))	

}
func Testgetuser(t *testing.T) {
    req, err := http.NewRequest("GET", "/getuser", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(getuser)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }else{
		fmt.Print("No error")
	}
	
    
}
func Getposts(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type","application/json")
	Clientt, _ := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.9w0dh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := Clientt.Database("admin").Collection("admin")
	fmt.Print(collection.FindOne(context.TODO(),request.URL.String()))	

}
func TestGetposts(t *testing.T) {
    req, err := http.NewRequest("GET", "/Getposts", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Getposts)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }else{
		fmt.Print("No error")
	}
	
    
}