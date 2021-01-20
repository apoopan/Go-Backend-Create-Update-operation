package account

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository //inetrface in db
	logger     log.Logger //to log whats going on
}

//NewService is ...
func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateApp(ctx context.Context, environment string, version int, appname string) (string, error) {
	logger := log.With(s.logger, "method", "CreateApp")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	app := App{
		id:          id,
		environment: environment,
		version:     version,
		appname:     appname,
	}
	if err := s.repository.CreateApp(ctx, app); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Created app", id)
	return "SUCCESS", nil
}

func (s service) GetUser(ctx context.Context) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetUser")
	var appname interface{}
	appname, err := s.repository.GetUser(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	fmt.Println(appname)
	return appname, nil
}

/*
func (s service) UpdateUser(ctx context.Context, id string, email string, password string, city string, age int) (string, error) {
	logger := log.With(s.logger, "method", "UpdateUser")

	user := User{
		ID:       id,
		Email:    email,
		Password: password,
		City:     city,
		Age:      age,
	}
	if err := s.repository.UpdateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Updated user", id)
	return "SUCCESS", nil

}
*/
