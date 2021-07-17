package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"net"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

// var dbInsert

type Payload struct {
	TS         string                 `json:"ts"`
	Sender     string                 `json:"sender"`
	SentFromIP string                 `json:"sent-from-ip"`
	Priority   int                    `json:"priority"`
	Message    map[string]interface{} `json:"message"`
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func createMessageHash(payload Payload) uint32 {
	// turning message back into a string (as how it came in from the handler)
	// concatenating all payload values then hashing
	messageString, _ := json.Marshal(payload.Message)
	return hash(payload.TS + payload.Sender + payload.SentFromIP + string(messageString))
}

func addPayloadToDB(payload Payload) {
	payloadHash := createMessageHash(payload)
	messageString, _ := json.Marshal(payload.Message)
	values := fmt.Sprintf("'%s','%s','%s','%s','%s'", fmt.Sprint(payloadHash), payload.TS, payload.Sender, string(messageString), payload.SentFromIP)
	query := fmt.Sprintf("INSERT INTO Payloads VALUES(%s)", values)
	fmt.Println("query: ", query)
	dbInsert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer dbInsert.Close()
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
		// adding valid payload to db
		addPayloadToDB(payload)
	} else {
		// returning error: invalid payload to sender
		http.Error(w, "error: invalid payload", http.StatusBadRequest)
		return
	}

	// debug response for success
	fmt.Fprintf(w, "Payload Hit")
}

func handleRequests(db *sql.DB) {
	// creating router with strict slash rule
	router := mux.NewRouter().StrictSlash(true)
	// setting up endpoint and handler for payload
	router.HandleFunc("/payload", handlePayload).Methods("POST")
	// setting up listener on port 8080 with router
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	// initializing sql server connection
	// using root for this demo but a proper implementation would use a created account with appropriate permissions
	var dbErr error
	db, dbErr = sql.Open("mysql", "root:rootpassword@tcp(127.0.0.1:3306)/mydb")
	if dbErr != nil {
		panic(dbErr.Error())
	}
	defer db.Close()

	// initializing web service
	handleRequests(db)
}
