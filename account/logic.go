package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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

func (s service) CreateApp(ctx context.Context, id string, environment string, version int, appname string) (string, error) {
	logger := log.With(s.logger, "method", "CreateApp")

	//uuid, _ := uuid.NewV4()
	//id := uuid.String()
	app := App{
		ID:          id,
		Environment: environment,
		Version:     version,
		Appname:     appname,
	}
	if err := s.repository.CreateApp(ctx, app); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Created app", id)
	return "SUCCESS", nil
}

func (s service) GetApp(ctx context.Context) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetApp")
	var environment interface{}
	environment, err := s.repository.GetApp(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return environment, nil
}

func (s service) UpdateApp(ctx context.Context, id string, environment string, version int, appname string) (string, error) {
	logger := log.With(s.logger, "method", "UpdateUser")
	app := App{
		ID:          id,
		Environment: environment,
		Version:     version,
		Appname:     appname,
	}

	if err := s.repository.UpdateApp(ctx, app); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Updated app", id)
	return "SUCCESS", nil

}
