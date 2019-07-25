package userdb

import (
	"context"
	"net/http"

	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/database"
	"github.com/PGo-Projects/tasky/internal/userdb/user"
	"github.com/PGo-Projects/webauth"
	"github.com/spf13/viper"
)

var (
	dbClient = database.MustMongoClient(context.TODO(), "mongodb://localhost:27017")
)

func GetUser(username string) (*user.User, error) {
	dbName := viper.GetString(config.DBName)
	users := dbClient.Database(dbName).Collection("users")

	filter := user.User{
		Username: username,
	}
	var u user.User
	err := users.FindOne(context.TODO(), filter).Decode(&u)
	return &u, err
}

func ReplaceUser(filter *user.User, user *user.User) error {
	dbName := viper.GetString(config.DBName)
	users := dbClient.Database(dbName).Collection("users")

	result := users.FindOneAndReplace(context.TODO(), filter, user)
	return result.Err()
}

func GenerateUserEntryHook(w http.ResponseWriter, r *http.Request, c webauth.Credentials) error {
	dbName := viper.GetString(config.DBName)
	users := dbClient.Database(dbName).Collection("users")

	u := user.User{
		Username:      c.Username,
		UnusedIndices: "1",
	}
	_, err := users.InsertOne(context.TODO(), u)
	return err
}
