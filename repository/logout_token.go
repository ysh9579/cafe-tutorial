package repository

import (
	"github.com/pkg/errors"
	"hello-cafe/internal/db"
	"hello-cafe/repository/dao"
)

type LogoutTokenRepository interface {
	Create(adminSeq int64, token string) error
	GetLogoutToken(adminSeq int64, token string) (*dao.LogoutToken, error)
}

type logoutTokenRepository struct{}

func NewLogoutTokenRepository() LogoutTokenRepository {
	return &logoutTokenRepository{}
}

func (r *logoutTokenRepository) Create(adminSeq int64, token string) error {
	t := &dao.LogoutToken{
		AdminSeq: adminSeq,
		Token:    token,
	}

	if err := db.Conn().Create(&t).Error; err != nil {
		return errors.Wrap(err, "failed to create logout")
	}

	return nil
}

func (r *logoutTokenRepository) GetLogoutToken(adminSeq int64, token string) (*dao.LogoutToken, error) {
	t := new(dao.LogoutToken)

	if err := db.Conn().
		Where("admin_seq = ?", adminSeq).
		Where("token = ?", token).
		Take(&t).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to get logout token by admin_seq(%d) and token (%s)", adminSeq, token)
	}

	return t, nil
}
