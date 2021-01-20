package account

import (
	"context"
)

//Service is ...
type Service interface {
	CreateApp(ctx context.Context, environment string, version int, appname string) (string, error)
	GetUser(ctx context.Context) (interface{}, error)
	//UpdateUser(ctx context.Context, id string, email string, password string, city string, age int) (string, error)
}
