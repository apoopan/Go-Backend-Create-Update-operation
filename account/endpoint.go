package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints is...
type Endpoints struct {
	CreateApp endpoint.Endpoint
	//GetUser    endpoint.Endpoint
	//UpdateUser endpoint.Endpoint
}

//MakeEndpoints is...
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateApp: makeCreateAppEndpoint(s),
		//GetUser:    makeGetUserEndpoint(s),
		//UpdateUser: makeUpdateUserEndpoint(s),
	}
}

func makeCreateAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAppRequest)
		ok, err := s.CreateApp(ctx, req.environment, req.version, req.appname)
		return CreateAppResponse{Ok: ok}, err
	}
}

/*
func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		data, err := s.GetUser(ctx)
		return GetUserResponse{Data: data, Err: err}, nil
	}

}

func makeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		ok, err := s.UpdateUser(ctx, req.ID, req.Email, req.Password, req.City, req.Age)
		return UpdateUserResponse{Ok: ok}, err
	}
}
*/
