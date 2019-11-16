package repo_impl

import (
	"context"
	"github.com/sangnguyen09/go_template/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/sangnguyen09/go_template/db/mongo"
	"github.com/sangnguyen09/go_template/repository"
	"github.com/sangnguyen09/go_template/models"
	//"github.com/sangnguyen09/go_template/helpers"
	"github.com/sangnguyen09/go_template/config"
)

type UserRepoImpl struct {
	mongo	*mongo.Mongo
}
func NewUserRepo(client *mongo.Mongo) repository.UserRespository {
	return &UserRepoImpl{
		mongo: client,
	}
}
func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq models.LoginRequest) (models.User, error) {

	var user models.User
	filter := bson.M{"username": loginReq.Username, "password": loginReq.Password}

	err := u.mongo.Client.Database(config.Config.Mongo.DatabaseName).Collection("users").FindOne(context, filter).Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepoImpl) Register(context context.Context,user models.User)(string, error)  {
	newID, _ := helpers.GetNextID(u.mongo.Client.Database(config.Config.Mongo.DatabaseName).Collection("counter"),
		"user_id_seq")
	user.UserId= newID
	_, err := u.mongo.Client.Database(config.Config.Mongo.DatabaseName).Collection("users").InsertOne(context, user)
	if err != nil {
		return "Lỗi", err
	}
	return "OK", nil
}
func (u *UserRepoImpl) CheckExist(context context.Context,username string) bool {
	var user models.User
	filter := bson.M{"username": username}
	 err := u.mongo.Client.Database(config.Config.Mongo.DatabaseName).Collection("users").FindOne(context, filter).Decode(&user)
	if err != nil {
		return false // không tồn tại
	}
	return true // tồn tại
}
