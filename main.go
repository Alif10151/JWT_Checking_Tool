package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CheckReq struct {
	JWT    string `json:"jwt"`
	Secret string `json:"secret"`
}

type ReturnResponse struct {
	Header         json.RawMessage `json:"header,omitempty"`
	Payload        json.RawMessage `json:"payload,omitempty"`
	ValidSignature bool            `json:"valid_signature"`
	Error          string          `json:"error,omitempty"`
}

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
