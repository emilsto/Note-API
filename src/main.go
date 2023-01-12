package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":5000"

func main() {
	router := mux.NewRouter()

	//endpoints, handlers, and http method
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/notes", getAllNotes).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/notes/{id}", getNote).Methods("GET")

	// Start and listen on port 5000
	//print success message
	fmt.Println("starting server at port 5000")
	http.ListenAndServe(port, router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthy 🔋⚡"))
}
