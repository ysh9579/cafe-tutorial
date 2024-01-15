package dao

import (
	"time"

	"hello-cafe/internal/apierror"
	"hello-cafe/internal/strcheck"
)

type Admin struct {
	AdminSeq int64     `gorm:"Column:admin_seq;PRIMARY_KEY"`
	Phone    string    `gorm:"Column:phone"`
	Password string    `gorm:"Column:password"`
	Name     string    `gorm:"Column:name"`
	RegDT    time.Time `gorm:"Column:reg_dt"`
	ModDT    time.Time `gorm:"Column:mod_dt"`
}

func (a Admin) TableName() string {
	return "admin"
}

func NewAdmin(phone, password, name string) (*Admin, error) {
	if !strcheck.ValidatePhone(phone) {
		return nil, apierror.ErrInvalidPhone
	}

	if !strcheck.ValidatePassword(password) {
		return nil, apierror.ErrInvalidPassword
	}

	return &Admin{
		Phone:    phone,
		Password: password,
		Name:     name,
		RegDT:    time.Now(),
		ModDT:    time.Now(),
	}, nil
}
