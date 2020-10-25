package main

import (
	Controller "KlikDigital/controller"
	"KlikDigital/utils"
	"fmt"
	"github.com/gorilla/mux"
)
func main(){
	mainRouter, _ := RegisterRoutes()
	port := "12345"

	fmt.Println("Server is listening  to port", port)

	utils.Listen(":"+port, mainRouter)
}

func RegisterRoutes() (*mux.Router, error){
	r := mux.NewRouter()

	r.HandleFunc("/users/register", Controller.Register).Methods("POST")
	r.HandleFunc("/create/transaction", Controller.CreateTransaction).Methods("POST")
	r.HandleFunc("/users/login", Controller.Login).Methods("POST")
	return r, nil
}
