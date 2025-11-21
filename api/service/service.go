package service

import (
	"fmt"

	"github.com/ingenziart/myapp/api/dto"
	"github.com/ingenziart/myapp/db"
	"github.com/ingenziart/myapp/models"
)

// creting new user
func CreateUser(createDto dto.CreateUserDto) (*models.User, error) {
	user := models.User{
		FullName:     createDto.FullName,
		Email:        createDto.Email,
		Phone:        createDto.Phone,
		PasswordHash: createDto.PasswordHash,
		Role:         createDto.Role,
		Status:       createDto.Status,
	}
	fmt.Print("created user ", createDto)

	//save to db

	if err := db.DB.Create(&user).Error; err != nil {
		return nil, err

	}
	return &user, nil

}

// get use with id
func GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, "id= ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id string, updateDto dto.UpdateUserDto) (*models.User, error) {

	var user models.User
	//check the existance
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err

	}
	//updating

	updates := map[string]interface{}{}
	if updateDto.FullName != nil {
		updates["fullName"] = *updateDto.FullName

	}
	if updateDto.Phone != nil {
		updates["phone"] = *updateDto.Phone
	}
	//save to db
	if len(updates) > 0 {
		if err := db.DB.Model(&user).Updates(updates).Error; err != nil {
			return nil, err
		}

	}
	return &user, nil

} // professiona way using model with updates not save it only change the map you created .

func UpdateStatus(id string, StatusDto dto.UpdateStatusDTO) (*models.User, error) {
	// finfing id
	var user models.User

	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	//update status
	StatusUpdate := map[string]interface{}{}

	if StatusDto.Status != nil {
		StatusUpdate["status"] = *StatusDto.Status

	}
	//save to db

	if len(StatusUpdate) > 0 {
		if err := db.DB.Model(&user).Updates(StatusUpdate).Error; err != nil {
			return nil, err
		}
	}
	return &user, nil

}
