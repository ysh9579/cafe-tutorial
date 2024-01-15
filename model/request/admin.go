package request

import (
	"hello-cafe/internal/apierror"
	"hello-cafe/internal/strcheck"
	"hello-cafe/internal/valid"
)

type Admin struct {
	Phone    *string `json:"phone,omitempty"`
	Password *string `json:"password,omitempty"`
	Name     string  `json:"name,omitempty"`
}

func (a *Admin) Validate() error {
	switch {
	case valid.IsNil(a.Phone):
		return apierror.ErrNilPhone
	case valid.IsNil(a.Password):
		return apierror.ErrNilPassword
	}

	if !strcheck.ValidatePhone(*a.Phone) {
		return apierror.ErrInvalidPhone
	}

	if !strcheck.ValidatePassword(*a.Password) {
		return apierror.ErrInvalidPassword
	}

	return nil
}

type SignOut struct {
	Phone *string `json:"phone"`
	Token *string `json:"token"`
}

func (l *SignOut) Validate() error {
	switch {
	case valid.IsNil(l.Phone):
		return apierror.ErrNilPhone
	case valid.IsNil(l.Token):
		return apierror.ErrNilAccessToken
	}

	if !strcheck.ValidatePhone(*l.Phone) {
		return apierror.ErrInvalidPhone
	}

	return nil
}
