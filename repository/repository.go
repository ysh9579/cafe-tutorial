package repository

import (
	"github.com/pkg/errors"
	"hello-cafe/internal/valid"
)

type Repository interface {
	Admin() AdminRepository
	Item() ItemRepository
	Logout() LogoutTokenRepository
}

type repository struct {
	admin  AdminRepository
	item   ItemRepository
	logout LogoutTokenRepository
}

func (r *repository) Validate() error {
	switch {
	case valid.IsNil(r.admin):
		return errors.New("admin repository is nil")
	case valid.IsNil(r.item):
		return errors.New("item repository is nil")
	case valid.IsNil(r.logout):
		return errors.New("logout token repository is nil")
	}

	return nil
}

func NewRepository() (Repository, error) {
	r := &repository{
		admin:  NewAdminRepository(),
		item:   NewItemRepository(),
		logout: NewLogoutTokenRepository(),
	}

	if err := r.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to create repository")
	}

	return r, nil
}

func (r *repository) Admin() AdminRepository {
	return r.admin
}

func (r *repository) Item() ItemRepository {
	return r.item
}

func (r *repository) Logout() LogoutTokenRepository {
	return r.logout
}
