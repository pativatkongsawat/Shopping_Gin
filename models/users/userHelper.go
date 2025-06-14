package users

import (
	"errors"
	"go_gin/helper"
	"time"

	"gorm.io/gorm"
)

type UserModelHelper struct {
	DB *gorm.DB
}

func (u *UserModelHelper) GetAllUser() ([]Users, error) {
	users := []Users{}
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserModelHelper) GetUser(fname, lname, email string, limit, page int) ([]Users, int64, error) {

	users := []Users{}

	var count int64

	offset := (page - 1) * limit

	if err := u.DB.Debug().Where("firstname LIKE ?", "%"+fname+"%", "lastname LIKE ?", "%"+lname+"%", "email LIKE ?", "%"+email+"%").
		Limit(limit).Offset(offset).
		Find(&users).Error; err != nil {

		return nil, 0, err
	}

	if err := u.DB.Debug().Model(&users).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (u *UserModelHelper) InsertUser(data []UsersInsert) ([]Users, error) {
	var usersToInsert []Users
	now := time.Now()

	for _, d := range data {
		user := Users{
			ID:        helper.GenerateUUID(),
			Firstname: d.Firstname,
			Lastname:  d.Lastname,
			Address:   d.Address,
			Email:     d.Email,
			Password:  helper.HashPassword(d.Password),
			CreatedAt: &now,
			UpdatedAt: &now,
			DeletedAt: nil,
		}
		usersToInsert = append(usersToInsert, user)
	}

	if err := u.DB.Create(&usersToInsert).Error; err != nil {
		return nil, err
	}

	return usersToInsert, nil
}

func (u *UserModelHelper) Register(data []Users) ([]Users, error) {

	var emails []string
	for _, user := range data {
		emails = append(emails, user.Email)
	}

	var existingUsers []Users
	if err := u.DB.Where("email IN ?", emails).Find(&existingUsers).Error; err != nil {
		return nil, err
	}

	if len(existingUsers) > 0 {
		return nil, errors.New("some emails already exist")
	}

	tx := u.DB.Begin()

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return data, nil
}
