package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"gopkg.in/mgo.v2"
)

//ErrRepo is ...
var ErrRepo = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *mgo.Database
	logger log.Logger
}

//NewRepo is ...
func NewRepo(db *mgo.Database, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "mongodb"),
	}, nil
}

func (repo *repo) CreateApp(ctx context.Context, app App) error {

	err := repo.db.C("configupdate").Insert(app)
	if err != nil {
		fmt.Println("Error occured inside CreateApp in repo")
		return err
	}
	fmt.Println("App Created:", app.environment, app.version, app.appname)

	return nil
}

/*
func (repo *repo) GetUser(ctx context.Context) (interface{}, error) {
	coll := repo.db.C("bloguser")
	data := []User{}
	err := coll.Find(bson.M{}).All(&data)
	if err != nil {
		fmt.Println("Error occured inside GetCUstomerById in repo")
		return "", err
	}
	return data, nil
}

func (repo *repo) UpdateUser(ctx context.Context, user User) error {
	f := bson.M{"id": user.ID}
	change := bson.M{"$set": bson.M{"password": user.Password, "email": user.Email, "city": user.City, "age": user.Age}}
	err := repo.db.C("bloguser").Update(f, change)

	if err != nil {
		return err
	}
	return nil
}
*/
