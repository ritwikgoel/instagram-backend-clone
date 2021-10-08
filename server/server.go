package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Print("initialization")
	http.ListenAndServe(":8081",nil)//nil is default handler
}