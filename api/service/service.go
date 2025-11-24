package service

import (
	"fmt"
	"math"

	"github.com/ingenziart/myapp/api/dto"
	"github.com/ingenziart/myapp/db"
	"github.com/ingenziart/myapp/models"
	customErr "github.com/ingenziart/myapp/utils/errors"
	"github.com/ingenziart/myapp/utils/pagination"
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
		Status:       models.Status(createDto.Status),
	}
	fmt.Print("created user ", createDto)

	//save to db

	if err := db.DB.Create(&user).Error; err != nil {
		return nil, err

	}
	return &user, nil

}
func GetAllUser() ([]models.User, error) {
	var user []models.User

	if err := db.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
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

func UpdateUserStatus(id string, dto dto.UpdateStatusDTO) (*models.User, error) {
	//chech the current status
	var user models.User

	if dto.Status == nil {
		return nil, fmt.Errorf("status required")
	}

	//check id to update
	if err := db.DB.First(&user, "id= ?", id).Error; err != nil {
		return nil, err
	}
	//validate new status

	newStatus := models.Status(*dto.Status)

	if !validation.IsValidateStatus(newStatus) {
		return nil, fmt.Errorf("invalid status ")
	}

	//save to db
	updates := map[string]interface{}{
		"status": newStatus,
	}

	//save to db
	if len(updates) > 0 {
		if err := db.DB.Model(&user).Updates(updates).Error; err != nil {
			return nil, err
		}

	}
	return &user, nil

}

// pagination
func FindAllUser(page, limit int) (*pagination.PaginationResponse, error) {

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	var user []models.User
	var total int64
	//query db without deleted data
	query := db.DB.Model(&models.User{}).Where("status != ?", models.StatusDeleted)

	//count toatal

	if err := query.Count(&total).Error; err != nil {
		return nil, err

	}
	//ofset to set paginanaion

	offset := (page - 1) * limit

	//get pagination data
	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&user).Error; err != nil {
		return nil, err
	}

	//
	lastPage := int(math.Ceil(float64(total) / float64(limit)))

	var prevPage, nextPage *int

	if page > 1 {
		prev := page - 1
		prevPage = &prev
	}
	if page < lastPage {
		next := page + 1
		nextPage = &next
	}
	return &pagination.PaginationResponse{
		Status:       "success",
		List:         user,
		Total:        total,
		PreviousPage: prevPage,
		NextPage:     nextPage,
		CurrentPage:  page,
		LastPage:     lastPage,
	}, nil

}

// soft delete
func SoftDeleteUser(id string) error {
	//check in db
	var user models.User
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		return err
	}
	//update status
	tx := db.DB.Begin()
	if err := tx.Model(&user).Update("status", models.StatusDeleted).Error; err != nil {
		tx.Rollback()
		return err

	}
	//soft delete
	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}
