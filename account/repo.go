package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
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

func (repo *repo) CreateUser(ctx context.Context, user User) error {

	err := repo.db.C("bloguser").Insert(user)
	if err != nil {
		fmt.Println("Error occured inside CreateUser in repo")
		return err
	}
	fmt.Println("User Created:", user.Email, user.Password, user.City, user.Age)

	return nil
}

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
