package service

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"hello-cafe/internal/apierror"
	"hello-cafe/internal/internaljwt"
	"hello-cafe/internal/valid"
	"hello-cafe/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	SignIn(phone string, password string) (accessToken string, err error)
	SignUp(phone string, password string, name string) error
	SignOut(phone string, token string) error
}

type adminService struct {
	repo repository.Repository
}

func NewAdminService(repo repository.Repository) (AdminService, error) {
	if valid.IsNil(repo) {
		return nil, errors.New("repository is nil")
	}

	return &adminService{repo: repo}, nil
}

func (s *adminService) SignIn(phone string, password string) (accessToken string, err error) {
	switch {
	case phone == "":
		return "", apierror.ErrNilPhone
	case password == "":
		return "", apierror.ErrNilPassword
	}

	admin, err := s.repo.Admin().GetAdminByPhone(phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.Wrap(err, "failed to get admin")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", apierror.ErrIDNotExist
	}

	if !s.checkPasswordHash(admin.Password, password) {
		return "", apierror.ErrIncorrectPassword
	}

	// 토큰 발행
	if accessToken, err = internaljwt.CreateJWT(phone); err != nil {
		return "", errors.Wrap(err, "failed to create jwt token")
	}

	return
}

func (s *adminService) checkPasswordHash(hashVal, userPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))
	if err != nil {
		return false
	} else {
		return true
	}
}

func (s *adminService) SignUp(phone string, password string, name string) error {
	switch {
	case phone == "":
		return apierror.ErrNilPhone
	case password == "":
		return apierror.ErrNilPassword
	}

	admin, err := s.repo.Admin().GetAdminByPhone(phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "failed to get admin")
	}

	if !valid.IsNil(admin) && admin.AdminSeq > 0 {
		return apierror.ErrDuplicatedAdmin
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to encrypt password")
	}

	encryptedPwd := string(bytes)
	if err := s.repo.Admin().Create(phone, encryptedPwd, name); err != nil {
		return errors.Wrap(err, "failed to sign up admin")
	}

	return nil
}

func (s *adminService) SignOut(phone string, token string) error {
	switch {
	case phone == "":
		return apierror.ErrNilPhone
	case token == "":
		return apierror.ErrNilAccessToken
	}

	admin, err := s.repo.Admin().GetAdminByPhone(phone)
	if err != nil {
		return errors.Wrap(err, "failed to get admin")
	}

	logoutToken, err := s.repo.Logout().GetLogoutToken(admin.AdminSeq, token)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "failed to get logout token")
	}

	if !valid.IsNil(logoutToken) {
		return apierror.ErrAlreadyLogout.SetInternal(fmt.Errorf("token_seq(%d) is already expired", logoutToken.TokenSeq))
	}

	if err := s.repo.Logout().Create(admin.AdminSeq, token); err != nil {
		return errors.Wrap(err, "failed to create logout token")
	}

	return nil
}
