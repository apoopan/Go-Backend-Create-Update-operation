package account

import (
	"context"
)

//Service is ...
type Service interface {
	CreateApp(ctx context.Context, id string, environment string, version int, appname string) (string, error)
	GetApp(ctx context.Context) (interface{}, error)
	UpdateApp(ctx context.Context, id string, environment string, version int, appname string) (string, error)
}
