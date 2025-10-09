package handlers

import "net/http"

func ManageCors(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false

}
