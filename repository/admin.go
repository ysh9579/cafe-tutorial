package repository

import (
	"github.com/pkg/errors"
	"hello-cafe/internal/db"
	"hello-cafe/repository/dao"
)

type AdminRepository interface {
	Create(phone, password, name string) error
	Get(adminSeq int64) (*dao.Admin, error)
	GetAdminByPhone(phone string) (*dao.Admin, error)
}

type adminRepository struct{}

func NewAdminRepository() AdminRepository {
	return &adminRepository{}
}

func (r *adminRepository) Create(phone, password, name string) error {
	admin, err := dao.NewAdmin(phone, password, name)
	if err != nil {
		return errors.Wrap(err, "failed to create new admin")
	}

	if err := db.Conn().Create(&admin).Error; err != nil {
		return errors.Wrap(err, "failed to create admin")
	}

	return nil
}

func (r *adminRepository) Get(adminSeq int64) (*dao.Admin, error) {
	admin := new(dao.Admin)
	if err := db.Conn().First(&admin, adminSeq).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get admin by admin sequence")
	}

	return admin, nil
}

func (r *adminRepository) GetAdminByPhone(phone string) (*dao.Admin, error) {
	admin := new(dao.Admin)

	if err := db.Conn().Where("phone = ?", phone).Take(&admin).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to get admin by phone(%s)", phone)
	}

	return admin, nil
}
