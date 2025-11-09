package main

import (
	"fmt"
	"net/http"
)

func ApiMainHandlerFirst (res http.ResponseWriter, req *http.Request) {
	fmt.Println("ApiMainHandler FIRST")
}

func ApiMainHandlerSecond (res http.ResponseWriter, req *http.Request) {
	fmt.Println("ApiMainHandler SECOND")
}