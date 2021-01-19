package account

import (
	"context"
	"encoding/json"
	"net/http"
	//"github.com/gorilla/mux"
)

type (
	CreateAppRequest struct {
		environment string `json:"environment"`
		version     int    `json:"version"`
		appname     string `json:"app-name"`
	}

	CreateAppResponse struct {
		Ok string `json:"ok"`
	}

	/*
		//GetUserRequest is...
		GetUserRequest struct {
		}
		//GetUserResponse is ...
		GetUserResponse struct {
			Data interface{} `json:"user"`
			Err  error       `json:"error,omitempty"`
		}

		//UpdateUserRequest is ...
		UpdateUserRequest struct {
			ID       string `json:"id"`
			Email    string `json:"email"`
			Password string `json:"password"`
			City     string `json:"city"`
			Age      int    `json:"age"`
		}

		//UpdateUserResponse is...
		UpdateUserResponse struct {
			Ok string `json:"ok"`
		}*/
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

/*
func encodeUpdateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeUpdateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest

	return req, nil
} */
