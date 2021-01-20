package account

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints is...
type Endpoints struct {
	CreateApp endpoint.Endpoint
	GetApp    endpoint.Endpoint
	UpdateApp endpoint.Endpoint
}

//MakeEndpoints is...
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateApp: makeCreateAppEndpoint(s),
		GetApp:    makeGetAppEndpoint(s),
		UpdateApp: makeUpdateAppEndpoint(s),
	}
}

func makeCreateAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAppRequest)
		ok, err := s.CreateApp(ctx, req.ID, req.Environment, req.Version, req.Appname)
		return CreateAppResponse{Ok: ok}, err
	}
}

func makeGetAppEndpoint(s Service) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		data, err := s.GetApp(ctx)
		return GetAppResponse{Data: data, Err: err}, nil
	}

}

func makeUpdateAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAppRequest)
		ok, err := s.UpdateApp(ctx, req.ID, req.Environment, req.Version, req.Appname)
		return UpdateAppResponse{Ok: ok}, err
	}
}
