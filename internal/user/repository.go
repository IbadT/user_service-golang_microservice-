package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

type Repository interface {
	CreateUser(user User) error
	GetUser(userId uuid.UUID) (User, error)
	GetUserById(userId uuid.UUID) (User, error)
	ListUsers() ([]User, error)
	UpdateUser(user User) error
	DeleteUser(userId uuid.UUID) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) CreateUser(user User) error {
	return r.DB.Create(&user).Error
}

func (r *repository) GetUser(userId uuid.UUID) (User, error) {
	var user User
	err := r.DB.First(&user, "id = ?", userId).Error
	return user, err
}

func (r *repository) GetUserById(userId uuid.UUID) (User, error) {
	var user User
	err := r.DB.First(&user, "id = ?", userId).Error
	return user, err
}

func (r *repository) ListUsers() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *repository) UpdateUser(user User) error {
	return r.DB.Save(&user).Error
}

func (r *repository) DeleteUser(userId uuid.UUID) error {
	return r.DB.Delete(&User{}, userId).Error
}
