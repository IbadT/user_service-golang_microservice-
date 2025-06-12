package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID     uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title  string    `json:"title" gorm:"type:varchar(255);not null"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid"`
}

type User struct {
	gorm.Model
	Id    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Tasks []Task    `gorm:"foreignKey:UserID"`
}
