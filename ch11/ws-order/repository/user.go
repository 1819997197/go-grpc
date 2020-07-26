package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"go-grpc/ch11/ws-order/vars"
	"go.elastic.co/apm/module/apmgorm"
)

type User struct {
	Id   int64  `orm:"id"`
	Name string `orm:"name"`
}

func (u User) TableName() string {
	return "user"
}

type userModel struct {
}

func NewUserModel() *userModel {
	return &userModel{}
}

func (u *userModel) FindList(ctx context.Context) ([]*User, error) {
	list := []*User{}
	db := apmgorm.WithContext(ctx, vars.DB)
	err := db.Model(&User{}).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return list, nil
}
