package service

import (
	"errors"
	"math"
	"strings"

	"github.com/ingenziart/myapp/api/dto"
	"github.com/ingenziart/myapp/db"
	"github.com/ingenziart/myapp/models"
	"github.com/ingenziart/myapp/utils/pagination"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrEmailInUse    = errors.New("email already in use")
	ErrHashPassword  = errors.New("failed to hash password")
	ErrCreateUser    = errors.New("failed to create user")
	ErrStatus        = errors.New("invalid status")
	ErrRole          = errors.New("invalid role")
	ErrUserNotFound  = errors.New("user not found")
	ErrFieldToUpdate = errors.New("no fielad  to update ")
)

// creting new user
func CreateUser(createDto dto.CreateUserDto) (*models.User, error) {

	//normalize input for email and phone
	email := strings.TrimSpace(strings.ToLower(createDto.Email))
	phone := strings.TrimSpace(createDto.Phone)

	//validate status and role(emums )
	status := models.Status(createDto.Status)
	role := models.Role(createDto.Role)

	if !status.IsValid() {
		return nil, ErrStatus

	}

	if !role.IsValid() {
		return nil, ErrRole

	}

	// check if email exist
	var existing models.User

	if err := db.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil, ErrEmailInUse

	}
	//hash the pasword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, ErrHashPassword

	}

	user := models.User{
		FullName:     createDto.FullName,
		Email:        email,
		Phone:        phone,
		PasswordHash: string(hashedPassword),
		Role:         role,
		Status:       status,
	}

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound //return consistance error

		}
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id string, updateDto dto.UpdateUserDto) (*models.User, error) {

	var user models.User
	//check the existance
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err

	}

	//map all thing to be changed is like empty map (object in js )
	updates := map[string]interface{}{}

	//update info
	if updateDto.FullName != nil {
		updates["fullName"] = strings.TrimSpace(*updateDto.FullName)
	}

	if updateDto.Email != nil {

		//check if email is not used before
		var temp models.User
		email := strings.TrimSpace(strings.ToLower(*updateDto.Email))
		if err := db.DB.Where("email = ? AND id <> ?", email, id).First(&temp).Error; err == nil {
			return nil, ErrEmailInUse

		}
		updates["email"] = email

	}
	//password
	if updateDto.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*updateDto.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, ErrHashPassword
		}
		updates["password"] = string(hashed)

	}

	//update role
	if updateDto.Role != nil {
		role := models.Role(*updateDto.Role)

		if !role.IsValid() {
			return nil, ErrRole

		}
		updates["role"] = role

	}
	if len(updates) == 0 {
		return nil, ErrFieldToUpdate
	}
	//updates to db
	if err := db.DB.Model(&user).Updates(updates).Error; err != nil {
		return nil, err

	}
	return &user, nil

}

func UpdateUserStatus(id string, dto dto.UpdateStatusDTO) (*models.User, error) {

	var user models.User

	//check id to update
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err

	}
	if dto.Status == nil {
		return nil, ErrFieldToUpdate
	}
	status := models.Status(*dto.Status)

	if !status.IsValid() {
		return nil, ErrStatus

	}
	updateStatus := map[string]interface{}{
		"status": status,
	}

	if err := db.DB.Model(&user).Updates(updateStatus).Error; err != nil {
		return nil, err

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

// soft deleteFU
func SoftDeleteUser(id string) error {
	//check in db to get id
	var user models.User
	tx := db.DB.Begin()
	if err := tx.First(&user, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	//update in db
	if err := tx.Model(&user).Update("status", models.StatusDeleted).Error; err != nil {
		tx.Rollback()
		return err
	}
	//delete
	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return err

	}
	return tx.Commit().Error

}

//restore

func RestoreUser(id string) (*models.User, error) {
	//check in db even deleted one using unscope func
	var user models.User
	if err := db.DB.Unscoped().First(&user, "id = ?", id).Error; err != nil {
		return nil, err

	}
	//the field of gorm.deletedAT
	if !user.DeletedAt.Valid {
		return &user, nil
	}
	//remove deleteAt
	tx := db.DB.Begin()
	if err := tx.Model(&user).Update("deleted_at", nil).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	//update status

	if err := tx.Model(&user).Update("status", models.StatusActive).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &user, tx.Commit().Error

}
