package account

import "context"

type App struct {
	ID          string `json:"id"`
	Environment string `json:"environment"`
	Version     int    `json:"version"`
	Appname     string `json:"app-name"`
}

// Repository is ...
type Repository interface {
	CreateApp(ctx context.Context, app App) error
	GetApp(ctx context.Context) (interface{}, error)
	UpdateApp(ctx context.Context, app App) error
}

/*
{
    "id" : "version-config",
    "environment" : "dev",
    "version" : "3.1.6",
    "app-name" : "jiocinema"
}
*/
