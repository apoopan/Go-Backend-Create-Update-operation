package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	//"go.mongodb.org/mongo-driver/mongo"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	fmt.Println("App Created:", app.ID, app.Environment, app.Version, app.Appname)

	return nil
}

func (repo *repo) GetApp(ctx context.Context) (interface{}, error) {
	coll := repo.db.C("configupdate")
	data := []App{}
	err := coll.Find(bson.M{}).All(&data)
	if err != nil {
		fmt.Println("Error occured inside GetAppById in repo")
		return "", err
	}
	return data, nil
}

func (repo *repo) UpdateApp(ctx context.Context, app App) error {
	f := bson.M{"id": app.ID}
	change := bson.M{"$set": bson.M{"environment": app.Environment, "version": app.Version, "appname": app.Appname}}
	//err := repo.db.C("configupdate").UpdateOne(ctx, f, change)
	//opts := options.Update().SetUpsert(true)
	col := repo.db.C("configupdate")
	info, err := col.Upsert(f, change)
	fmt.Println(info)
	if err != nil {
		return err
	}
	return nil
}
