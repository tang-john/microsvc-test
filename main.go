package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/johntang/db-employee/handlers"
)

const Port = ":8080"


func main() {
	mux := httprouter.New()
	mux.GET("/", handlers.JamesHandler)
	mux.GET("/james", handlers.JamesHandler)
	mux.GET("/microsvc-test", handlers.MicrosvcTestHandler)


	port, ok := os.LookupEnv("PORT")
	if ok {
		port = ":" + port
	} else {
		port = Port
	}

	fmt.Println("Starting microservice listener on port :", port)
	http.ListenAndServe(port, mux)
}


func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
