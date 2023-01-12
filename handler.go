package main

import (
	"encoding/json"
	"net/http"
)

var users = []User{
	{ID: 1, Email: "test@example.com", Username: "test", Password: "test"},
	{ID: 2, Email: "test2@example.com", Username: "test2", Password: "test2"},
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
