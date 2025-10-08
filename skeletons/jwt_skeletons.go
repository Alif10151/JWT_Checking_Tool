package skeletons

import "encoding/json"

type CheckReq struct {
	JWT    string `json:"jwt"`
	Secret string `json:"secret"`
	Alg    string `json:"alg"`
}

type ReturnResponse struct {
	Header         json.RawMessage `json:"header,omitempty"`
	Payload        json.RawMessage `json:"payload,omitempty"`
	ValidSignature bool            `json:"valid_signature"`
	Error          []string        `json:"error,omitempty"`
}
