package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Payload struct {
	TS         string                 `json:"ts"`
	Sender     string                 `json:"sender"`
	SentFromIP string                 `json:"sent-from-ip"`
	Priority   int                    `json:"priority"`
	Message    map[string]interface{} `json:"message"`
}

func validateTimestamp(timestamp string) bool {
	// a unix timestamp is valid if it is an integer (seconds from unix epoch)
	// if string to int fails, it is not a valid integer
	if _, err := strconv.Atoi(timestamp); err != nil {
		return false
	} else {
		return true
	}
}

func validateSender(sender string) bool {
	// sender is of type string if it was successfully parsed
	// return false if there is no sender
	return sender != ""
}

func validateSentFromIP(ip string) bool {
	if ip != "" {
		if validIP := net.ParseIP(ip); validIP != nil {
			// validIP is not nil thus it was successfully parsed and is a valid IPv4 address
			return true
		} else {
			// validIP is nil thus it is an invalid IPv4 address
			return false
		}
	} else {
		// No IP is a valid state for the payload
		return true
	}
}

func validateMessage(message map[string]interface{}) bool {
	if len(message) > 0 {
		// at least one value in message
		return true
	} else {
		// no values in message or message wasn't present in json thus invalid
		return false
	}
}

func validatePayload(payload Payload) bool {
	return (validateTimestamp(payload.TS) && validateSender(payload.Sender) && validateSentFromIP(payload.SentFromIP) && validateMessage(payload.Message))
}

func handlePayload(w http.ResponseWriter, r *http.Request) {
	// instantiating payload struct
	var payload Payload

	// creating strict decoder that will throw error if an unknown field is included
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validating payload and either handling data or returning validation error
	if validatePayload(payload) {
		// TODO: handle queueing and storing (sql) data
	} else {
		// returning error: invalid payload to sender
		http.Error(w, "error: invalid payload", http.StatusBadRequest)
		return
	}

	// debug response for success
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
