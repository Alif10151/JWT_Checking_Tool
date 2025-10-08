package main

import (
	"JWT_Checking_Tool/handlers"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/decode", handlers.CreateJWT)
	fmt.Println("Server is running on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Unfortunately Error Occurs", err)
	}

}
