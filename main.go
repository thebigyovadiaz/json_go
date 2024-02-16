package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func decodeJSON(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	_, _ = fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
}

func encodeJSON(w http.ResponseWriter, r *http.Request) {
	peter := &User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}

	err := json.NewEncoder(w).Encode(peter)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/decode", decodeJSON)
	http.HandleFunc("/encode", encodeJSON)

	_ = http.ListenAndServe(":2323", nil)
}
