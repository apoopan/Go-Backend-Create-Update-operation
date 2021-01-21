package account

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//NewHTTPServer is ...
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/app").Handler(httptransport.NewServer(
		endpoints.CreateApp,
		decodeAppReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/getapp").Handler(httptransport.NewServer(
		endpoints.GetApp,
		decodeGetAppRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/update").Handler(httptransport.NewServer(
		endpoints.UpdateApp,
		decodeUpdateReq,
		encodeUpdateResponse,
	))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
