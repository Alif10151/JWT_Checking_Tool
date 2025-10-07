package main

import (
	"fmt"
	"net/http"
)

func startHendler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server Works")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/start", startHendler)
	fmt.Println("Server is running on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Unfortunately Error Occurs", err)
	}

}
