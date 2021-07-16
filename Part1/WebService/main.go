package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Payload struct {
	TS         string
	Sender     string
	SentFromIP string
	Priority   int
	Message    map[string]interface{}
}

func validatePayload(payload Payload) bool {

}

func handlePayload(w http.ResponseWriter, r *http.Request) {
	// decoding body of request and handling read/decode errors
	var payload Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validating payload and either handle data or error
	if validatePayload(payload) {
		// TODO: handle queueing and storing (sql) data
	} else {
		// returning error: invalid payload to sender
		http.Error(w, "error: invalid payload", http.StatusBadRequest)
	}
	fmt.Print("payload", payload)

	// fmt.Println("Post Recieved and decoded", payload, payload.Message)
	fmt.Fprintf(w, "Payload Hit")
}

func handleRequests() {
	// creating router with strict slash rule
	router := mux.NewRouter().StrictSlash(true)
	// setting up endpoint and handler for payload
	router.HandleFunc("/payload", handlePayload).Methods("POST")
	// setting up listener on port 8080 with router
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	// initializing web service
	handleRequests()
}
