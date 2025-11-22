package service

import (
	"fmt"

	"github.com/ingenziart/myapp/api/dto"
	"github.com/ingenziart/myapp/db"
	"github.com/ingenziart/myapp/models"
	customErr "github.com/ingenziart/myapp/utils/errors"
	"github.com/ingenziart/myapp/utils/validation"
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
	//check nill
	if updateDto.FullName == nil && updateDto.Phone == nil {
		return nil, customErr.ErrNoFieldToUpdate
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
			return nil, customErr.ErrNoFieldToUpdate
		}

	}
	return &user, nil

} // professiona way using model with updates not save it only change the map you created .

func UpdateStatus(id string, dto dto.UpdateStatusDTO) (*models.User, error) {
	//check id in db

	var user models.User

	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err

	}
	//check if there is status
	if dto.Status == nil {
		return nil, fmt.Errorf("status required")

	}
	//field to type  with condition
	NewStatus := models.Status(*dto.Status)

	//validate
	if !validation.IsValidateStatus(NewStatus) {
		return nil, fmt.Errorf("invalid response")

	}
	//update
	update := map[string]interface{}{
		"status": NewStatus,
	}
	if err := db.DB.Model(&user).Updates(update).Error; err != nil {
		return nil, err
	}
	return &user, nil

}
