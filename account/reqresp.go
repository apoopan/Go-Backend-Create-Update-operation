package account

import (
	"context"
	"encoding/json"
	"net/http"
	//"github.com/gorilla/mux"
)

type (
	CreateAppRequest struct {
		ID          string `json:"id"`
		Environment string `json:"environment"`
		Version     int    `json:"version"`
		Appname     string `json:"app-name"`
	}

	CreateAppResponse struct {
		Ok string `json:"ok"`
	}

	//GetUserRequest is...
	GetAppRequest struct {
	}
	//GetUserResponse is ...
	GetAppResponse struct {
		Data interface{} `json:"app"`
		Err  error       `json:"error,omitempty"`
	}

	//UpdateUserRequest is ...
	UpdateAppRequest struct {
		ID          string `json:"id"`
		Environment string `json:"environment"`
		Version     int    `json:"version"`
		Appname     string `json:"app-name"`
	}

	//UpdateUserResponse is...
	UpdateAppResponse struct {
		Ok string `json:"ok"`
	}
)

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeAppReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateAppRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeUpdateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeUpdateReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateAppRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetAppRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAppRequest

	return req, nil
}
