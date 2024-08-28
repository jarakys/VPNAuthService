package Managers

import (
	"VPNAuthService/DbModels"
	"VPNAuthService/Errors"
	"gorm.io/gorm"
	"time"
)

type AuthDbManager interface {
	Create() (*DbModels.UserModel, error)
	Update(user DbModels.UserModel) error
	Delete(userId string) error
	Get(userId string) (*DbModels.UserModel, error)
}

func NewAuthDbManager(db *gorm.DB) AuthDbManager {
	return &authDbManagerImpl{db: db}
}

type authDbManagerImpl struct {
	db *gorm.DB
}

func (a *authDbManagerImpl) Create() (*DbModels.UserModel, error) {
	userModel := DbModels.UserModel{
		LastVisit: time.Now().Unix(),
		IsPremium: false,
	}
	err := a.db.Create(&userModel).Error
	if err != nil {
		return nil, Errors.ErrorDbWrapperUtils(err)
	}
	return &userModel, nil
}

func (a *authDbManagerImpl) Update(user DbModels.UserModel) error {
	if err := a.db.Save(&user).Error; err != nil {
		return Errors.ErrorDbWrapperUtils(err)
	}
	return nil
}

func (a *authDbManagerImpl) Delete(userId string) error {
	if err := a.db.Delete(&DbModels.UserModel{}, "id = ?", userId).Error; err != nil {
		return Errors.ErrorDbWrapperUtils(err)
	}
	return nil
}

func (a *authDbManagerImpl) Get(userId string) (*DbModels.UserModel, error) {
	userModel := DbModels.UserModel{}
	if err := a.db.Where("id = ?", userId).First(&userModel).Error; err != nil {
		return nil, Errors.ErrorDbWrapperUtils(err)
	}
	return &userModel, nil
}
