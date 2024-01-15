package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hello-cafe/internal/apierror"
	"hello-cafe/model/request"
	"hello-cafe/model/response"
	"hello-cafe/service"
)

type ItemHandler interface {
	Create(ctx *gin.Context) // 상품 등록
	Update(ctx *gin.Context) // 상품 수정
	Delete(ctx *gin.Context) // 상품 삭제
	Find(ctx *gin.Context)   // 상품 리스트 조회
	Get(ctx *gin.Context)    // 상품 상세
	Search(ctx *gin.Context) // 상품 이름
}

type itemHandler struct {
	itemService service.ItemService
}

func NewItemHandler(itemService service.ItemService) (ItemHandler, error) {
	return &itemHandler{
		itemService: itemService,
	}, nil
}

func (h *itemHandler) Create(ctx *gin.Context) {
	req := request.CreateItem{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := h.itemService.Create(req); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	ctx.JSON(response.SimpleSuccess(http.StatusOK))
}

func (h *itemHandler) Update(ctx *gin.Context) {
	req := request.UpdateItem{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := h.itemService.Update(req); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	ctx.JSON(response.SimpleSuccess(http.StatusOK))
}

func (h *itemHandler) Delete(ctx *gin.Context) {
	req := request.DeleteItem{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := h.itemService.Delete(req.ItemSeq); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	ctx.JSON(response.SimpleSuccess(http.StatusOK))
}

func (h *itemHandler) Find(ctx *gin.Context) {
	queryAdminSeq := ctx.Query("admin_seq")
	if queryAdminSeq == "" {
		ctx.AbortWithStatusJSON(response.Failure(apierror.ErrInvalidAdmin))
		return
	}

	adminSeq, err := strconv.ParseInt(queryAdminSeq, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	queryLastItemSeq := ctx.DefaultQuery("last_item_seq", "0")

	lastItemSeq, err := strconv.ParseInt(queryLastItemSeq, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	queryLimit := ctx.DefaultQuery("limit", "10")

	limit, err := strconv.Atoi(queryLimit)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	items, err := h.itemService.Find(adminSeq, lastItemSeq, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	ctx.JSON(response.Success(items))
}

func (h *itemHandler) Get(ctx *gin.Context) {
	req := request.GetItem{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	item, err := h.itemService.Get(req.ItemSeq)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	ctx.JSON(response.Success(item))
}

func (h *itemHandler) Search(ctx *gin.Context) {
	queryAdminSeq := ctx.Query("admin_seq")
	if queryAdminSeq == "" {
		ctx.AbortWithStatusJSON(response.Failure(apierror.ErrInvalidAdmin))
		return
	}

	adminSeq, err := strconv.ParseInt(queryAdminSeq, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	queryText := ctx.Query("text")
	if queryText == "" {
		ctx.AbortWithStatusJSON(response.Failure(apierror.ErrNilSearchText))
		return
	}

	items, err := h.itemService.Search(adminSeq, queryText)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Failure(err))
		return
	}

	ctx.JSON(response.Success(items))
}
