package response

import (
	"net/http"

	"hello-cafe/internal/apierror"
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

const defaultSuccessMsg = "ok"

func SimpleSuccess(code int) (int, Response) {
	return code, Response{
		Meta: Meta{
			Code:    code,
			Message: defaultSuccessMsg,
		},
		Data: nil,
	}
}

func Success(data interface{}) (int, Response) {
	return http.StatusOK, Response{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: defaultSuccessMsg,
		},
		Data: data,
	}
}

func Failure(err error) (int, Response) {
	code := http.StatusInternalServerError
	msg := "서버 내부 오류 입니다."

	if apiErr, isAPIErr := apierror.IsAPIError(err); isAPIErr {
		code = apiErr.Code
		msg = apiErr.Message()
	}

	return code, Response{
		Meta: Meta{
			Code:    code,
			Message: msg,
		},
		Data: nil,
	}
}
