package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// ユーザのインターンフェース定義
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	//gorm.DBのインターフェースに依存させる
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// ユーザのリポジトリの実装
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	//emailをもとにユーザを取得する
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// 引数と返り値はインターフェースに依存させる
func (ur *userRepository) CreateUser(user *model.User) error {
	//ユーザを作成する
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
