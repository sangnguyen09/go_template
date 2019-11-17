package repo_impl

import (
	"context"
	"errors"
	"github.com/sangnguyen09/go_template/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (u *UserRepoImpl) Register(context context.Context,user models.User) error {
	newID, _ := helpers.GetNextID(u.mongo.Client.Database(config.Config.Mongo.DatabaseName).Collection("counter"),
		"user_id_seq")
	user.UserId= newID
	_, err := u.mongo.Client.Database(config.Config.Mongo.DatabaseName).Collection("users").InsertOne(context, user)
	if err != nil {
		return errors.New("Đăng kí thất bại")
	}
	return nil
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

func (u *UserRepoImpl) ComparePassword(context context.Context,pwdcurrent string, userId int) bool {
	var user models.User
	filter := bson.M{"password": pwdcurrent,"user_id":userId}
	err := u.mongo.Client.Database(config.Config.Mongo.DatabaseName).Collection("users").FindOne(context, filter).Decode(&user)
	if err != nil {
		return false // không trùng
	}
	return true // trùng mật khẩu
}

func (u *UserRepoImpl) UpdatePass(context context.Context,pwdnew string, userId int)  error {
	var user models.User
	filter := bson.M{"user_id":userId}
	update := bson.M{"$set":bson.M{"password":pwdnew}}
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(1) // 1 là tra ve tai lieu moi sau khi update, 0 là trước update
	err := u.mongo.Client.Database(config.Config.Mongo.DatabaseName).Collection("users").FindOneAndUpdate(context, filter, update, opts).Decode(&user)
	if err != nil {
		return  errors.New("Update thất bại")
	}

	return nil
}
func (u *UserRepoImpl) Delete(context context.Context, userId int)  error {
	var deleteDocument bson.M
	filter := bson.D{{"$and", []bson.D{ bson.D{{"user_id",userId}}, bson.D{{"user_id",bson.D{{"$ne",1}} }}, }}}// xoá user có id và id đó khác 1(admin)
	opts := options.FindOneAndDelete()
	opts.SetProjection(bson.D{{"username",1}}) // lấy ra field cần thiết
	err := u.mongo.Client.Database(config.Config.Mongo.DatabaseName).Collection("users").FindOneAndDelete(context, filter, opts).Decode(&deleteDocument)
	if err != nil {
		return  errors.New("User Không tồn tại")
	}

	return nil
}
