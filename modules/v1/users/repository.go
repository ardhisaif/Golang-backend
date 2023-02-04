package users

import (
	"errors"
	"fmt"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (*model.Users, error) {
	var user model.Users
	data := r.db.Find(&user)
	if data.Error != nil {
		return nil, data.Error
	}

	return &user, nil
}

func (r *repository) Register(user *model.User) (*model.User, error) {
	data := r.db.Create(&user)
	if data.Error != nil {
		return nil, data.Error
	}

	user.Password = ""

	return user, nil
}

func (r *repository) Update(user *model.User, id string) (*model.User, error) {
	
	fmt.Println(user)
	data := r.db.Model(&user).Where("user_id = ?",id).Updates(&user)
	if data.Error != nil {
		return nil, data.Error
	}
	fmt.Println(user)
	fmt.Println(data)

	u := r.db.Where("user_id = ? ", id).Limit(1).Find(&user)
	if u.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil 
}

func (r *repository) Delete(user *model.User, id string) (*model.User, error){
	err := r.db.Where("user_id = ?", id).Delete(&user).Error
	if err != nil {
		return nil, err
	}
	
	return user, nil
}