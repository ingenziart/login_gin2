package models

import (
	"time"

	"gorm.io/gorm"
)

// this is for validation
type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusDeleted  Status = "deleted"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusActive, StatusInactive, StatusDeleted:
		return true
	}
	return false
}

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
	RoleGuest Role = "guest"
)

func (s Role) IsValid() bool {
	switch s {
	case RoleAdmin, RoleUser, RoleGuest:
		return true
	}
	return false
}

type User struct {
	ID              string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	FullName        string         `gorm:"column:fullName;type:varchar(120);not null" json:"fullName"`
	Email           string         `gorm:"column:email;type:citext;unique;not null" json:"email"`
	Phone           string         `gorm:"column:phone;type:varchar(32);not null" json:"phone"`
	PasswordHash    string         `gorm:"column:password_hash;type:text;not null" json:"-"`
	Role            Role           `gorm:"column:role;type:varchar(32);not null" json:"role"`
	Status          Status         `gorm:"column:status;type:varchar(16);not null" json:"status"`
	EmailVerifiedAt *time.Time     `gorm:"column:email_verified_at;type:timestamptz" json:"emailVerified_at"`
	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime;not null" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;autoUpdateTime;not null" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"DeletedAt"`
	//Metadata        JSONB          `gorm:"type:jsonb;default:'{}'::jsonb"`
}
