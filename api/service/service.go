package service

import (
	"fmt"

	"github.com/ingenziart/myapp/api/dto"
	"github.com/ingenziart/myapp/db"
	"github.com/ingenziart/myapp/models"
)

func CreateUser(createDto dto.CreateUserDto) (*models.User, error) {
	use := models.User{
		FullName:     createDto.FullName,
		Email:        createDto.Email,
		Phone:        createDto.Phone,
		PasswordHash: createDto.PasswordHash,
		Role:         createDto.Role,
		Status:       createDto.Status,
	}
	fmt.Print("created user ", createDto)

	//save to db

	if err := db.DB.Create(&use).Error; err != nil {
		return nil, err

	}
	return &use, nil

}

func GetUserByID(id string) (*models.User, error) {
	var use models.User
	if err := db.DB.Where("id=?", id).First(&use); err != nil {
		return nil, err.Error
	}
	return &use, nil
}

func UpdateUserByID(updateDto dto.UpdateUserDto) (models.User, error) {
	use := models.User{}

}
