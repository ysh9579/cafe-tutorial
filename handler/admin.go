package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hello-cafe/model/request"
	"hello-cafe/model/response"
	"hello-cafe/service"
)

type AdminHandler interface {
	SignIn(ctx *gin.Context)  // 로그인
	SignUp(ctx *gin.Context)  // 회원가입
	SignOut(ctx *gin.Context) // 로그아웃
}

type adminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) (AdminHandler, error) {
	return &adminHandler{
		adminService: adminService,
	}, nil
}

func (h *adminHandler) SignIn(ctx *gin.Context) {
	req := request.Admin{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	token, err := h.adminService.SignIn(*req.Phone, *req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	ctx.JSON(response.Success(gin.H{"token": token}))
}

func (h *adminHandler) SignUp(ctx *gin.Context) {
	req := request.Admin{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := h.adminService.SignUp(*req.Phone, *req.Password, req.Name); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	ctx.JSON(response.SimpleSuccess(http.StatusOK))
}

func (h *adminHandler) SignOut(ctx *gin.Context) {
	req := request.SignOut{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := h.adminService.SignOut(*req.Phone, *req.Token); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	ctx.JSON(response.SimpleSuccess(http.StatusOK))
}
