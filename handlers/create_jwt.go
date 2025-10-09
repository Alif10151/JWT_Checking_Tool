package handlers

import (
	utility "JWT_Checking_Tool/Utility"
	"JWT_Checking_Tool/skeletons"
	"encoding/json"
	"net/http"
	"strings"
)

func CreateJWT(w http.ResponseWriter, r *http.Request) {
	if stop := ManageCors(w, r); stop {
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Post Method Accepted Only", http.StatusMethodNotAllowed)
		return
	}

	var req skeletons.CheckReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		http.Error(w, `{"error":["invalid json"]}`, http.StatusBadRequest)
		return
	}

	if req.JWT == "" || req.Secret == "" || req.Alg == "" {
		http.Error(w, `{"error":["jwt,secret and alg are neeeded"]}`, http.StatusBadRequest)
		return
	}

	if strings.ToUpper(req.Alg) != "HS256" {
		http.Error(w, `{"error":["HS256 Supports Only"]}`, http.StatusBadRequest)
		return
	}

	parts := strings.Split(req.JWT, ".")
	if len(parts) != 3 {
		response := skeletons.ReturnResponse{
			ValidSignature: false,
			Error:          []string{"Wrong JWT Format"}}
		json.NewEncoder(w).Encode(response)
		return
	}

	//parts[0]==header, parts[1]=payload, parts[2]=signature

	headerBytes, err := utility.DecodeB64Url(parts[0])
	if err != nil {
		response := skeletons.ReturnResponse{
			ValidSignature: false,
			Error:          []string{"Error In Header Decode"}}
		json.NewEncoder(w).Encode(response)
		return
	}

	payloadBytes, err := utility.DecodeB64Url(parts[1])
	if err != nil {
		response := skeletons.ReturnResponse{
			ValidSignature: false,
			Error:          []string{"Error In Payload Decode"}}
		json.NewEncoder(w).Encode(response)
		return
	}

	_, sigErr := utility.DecodeB64Url(parts[2])
	if sigErr != nil {
		response := skeletons.ReturnResponse{
			ValidSignature: false,
			Error:          []string{"Error In Signature Decode"}}
		json.NewEncoder(w).Encode(response)
		return
	}

	valid := utility.VerifyHS(parts[0], parts[1], parts[2], []byte(req.Secret))

	if !valid {
		response := skeletons.ReturnResponse{
			ValidSignature: false,
			Error:          []string{"Signature or secret key does not match"}}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := skeletons.ReturnResponse{
		Header:         headerBytes,
		Payload:        payloadBytes,
		ValidSignature: valid,
	}

	json.NewEncoder(w).Encode(response)

}
