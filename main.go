package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//endpoints, handlers, and http method
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/users", getAllUsers).Methods("GET")

	// Start and listen on port 5000
	http.ListenAndServe(":5000", router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthy 🔋⚡"))
}
