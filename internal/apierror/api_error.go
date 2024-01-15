package apierror

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNilPhone           = NewAPIError(http.StatusBadRequest, "핸드폰 번호를 입력해 주세요.")
	ErrNilPassword        = NewAPIError(http.StatusBadRequest, "비밀번호를 입력해 주세요")
	ErrInvalidPhone       = NewAPIError(http.StatusBadRequest, "핸드폰 번호 형식이 잘못 되었습니다.")
	ErrInvalidPassword    = NewAPIError(http.StatusBadRequest, "비밀번호 형식이 잘못 되었습니다.")
	ErrNilCategory        = NewAPIError(http.StatusBadRequest, "카테고리를 입력해 주세요.")
	ErrInvalidCategory    = NewAPIError(http.StatusBadRequest, "카테고리 형식이 잘못 되었습니다.")
	ErrNilBarcode         = NewAPIError(http.StatusBadRequest, "바코드를 입력해 주세요.")
	ErrNilPrice           = NewAPIError(http.StatusBadRequest, "가격을 입력해 주세요.")
	ErrNilCost            = NewAPIError(http.StatusBadRequest, "원가를 입력해 주세요.")
	ErrNilName            = NewAPIError(http.StatusBadRequest, "이름을 입력해 주세요.")
	ErrNilDescription     = NewAPIError(http.StatusBadRequest, "설명을 입력해 주세요.")
	ErrNilExpireDT        = NewAPIError(http.StatusBadRequest, "유통기한을 입력해 주세요.")
	ErrNilSize            = NewAPIError(http.StatusBadRequest, "사이즈를 입력해 주세요.")
	ErrNilSearchText      = NewAPIError(http.StatusBadRequest, "검색어를 입력해 주세요.")
	ErrInvalidSize        = NewAPIError(http.StatusBadRequest, "사이즈 형식이 잘못 되었습니다.")
	ErrDuplicatedAdmin    = NewAPIError(http.StatusBadRequest, "중복된 계정입니다.")
	ErrNilAccessToken     = NewAPIError(http.StatusBadRequest, "access token 정보를 입력해 주세요.")
	ErrInvalidAdmin       = NewAPIError(http.StatusBadRequest, "관리자 정보가 잘못 되었습니다.")
	ErrInvalidItem        = NewAPIError(http.StatusBadRequest, "상품 정보가 잘못 되었습니다.")
	ErrDuplicatedItem     = NewAPIError(http.StatusBadRequest, "중복된 상품입니다.")
	ErrNotExistItem       = NewAPIError(http.StatusBadRequest, "존재하지 않는 상품입니다.")
	ErrInvalidAccessToken = NewAPIError(http.StatusBadRequest, "엑세스 토큰 인증 실패.")
)

var (
	ErrIDNotExist        = NewAPIError(http.StatusUnauthorized, "존재하지 않는 계정입니다.")
	ErrIncorrectPassword = NewAPIError(http.StatusUnauthorized, "비밀번호가 잘못 되었습니다.")
	ErrAlreadyLogout     = NewAPIError(http.StatusUnauthorized, "이미 로그아웃 되었습니다.")
)

type APIError struct {
	Code     int
	Msg      string
	Internal error
}

func NewAPIError(code int, msg string) *APIError {
	return &APIError{
		Code: code,
		Msg:  msg,
	}
}

func IsAPIError(err error) (apiErr *APIError, ok bool) {
	if err == nil {
		return
	}

	ok = errors.As(err, &apiErr)

	return
}

func (e *APIError) SetInternal(err error) *APIError {
	if err != nil {
		e.Internal = err
	}
	return e
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d: %v", e.Code, e.Internal)
}

func (e *APIError) Message() string {
	return e.Msg
}
