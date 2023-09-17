package server

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendPostRequest(url string, payload any) (*http.Response, error) {
	// convert struct into json string
	jsonData, _ := json.MarshalIndent(payload, "", "\t")

	// call the auth service
	request, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	return http.DefaultClient.Do(request)
}

func ParseReponseBody(response *http.Response) (*jsonResponse, error) {
	var jsonRes jsonResponse
	err := json.NewDecoder(response.Body).Decode(&jsonRes)
	if err != nil {
		return nil, err
	}

	if jsonRes.Error {
		return nil, err
	}

	return &jsonRes, nil
}
