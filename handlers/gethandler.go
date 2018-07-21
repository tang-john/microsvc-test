package handlers

import (
	"encoding/json"
  "fmt"
	"log"
	"net/http"


	"github.com/julienschmidt/httprouter"
)

type person struct {
	Fname string
	Lname string
	Items []string
}


// JamesHandler will handle GET request to /james. This is used as a test
func JamesHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}


func MicrosvcTestHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Hello World")
}
