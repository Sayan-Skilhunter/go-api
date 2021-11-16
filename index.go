package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Payload struct {
	data string
}

var payload Payload

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user", getUsers).Methods("GET")
	http.ListenAndServe(":8080", r)
	log.Println("server is running in port 8080 ")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(payload)
	res := make(map[string]interface{})
	res["response"] = "Data Posted"
	des, _ := json.Marshal(res)
	w.Write([]byte(des))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(payload)

	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(payload)
	res := make(map[string]interface{})
	res["response"] = "Hit Get endpoint"
	des, _ := json.Marshal(res)
	w.Write([]byte(des))
}
