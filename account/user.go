package account

import "context"

type App struct {
	id          string `json:"id"`
	environment string `json:"environment"`
	version     int    `json:"version"`
	appname     string `json:"app-name"`
}

// Repository is ...
type Repository interface {
	CreateApp(ctx context.Context, app App) error
	//GetUser(ctx context.Context) (interface{}, error)
	//UpdateUser(ctx context.Context, user User) error
}

/*
{
    "id" : "version-config",
    "environment" : "dev",
    "version" : "3.1.6",
    "app-name" : "jiocinema"
}
*/
