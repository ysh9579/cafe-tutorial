package repository

import (
	"strings"
	"time"

	"github.com/LoperLee/golang-hangul-toolkit/hangul"
	"github.com/pkg/errors"
	"hello-cafe/internal/valid"
	"hello-cafe/model/request"
	"hello-cafe/repository/dao"
)

type ColumnMap interface {
	ToMap() map[string]interface{}
}

type UpdateItem struct {
	Category    *dao.ItemCategory
	Barcode     *string
	Price       *int64
	Cost        *int64
	Name        *string
	Description *string
	ExpireDT    *time.Time
	Size        *dao.ItemSize
}

func NewUpdateItem(r request.UpdateItem) (*UpdateItem, error) {
	if err := r.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate update item request")
	}

	item := &UpdateItem{
		Barcode:     r.Barcode,
		Price:       r.Price,
		Cost:        r.Cost,
		Name:        r.Name,
		Description: r.Description,
		ExpireDT:    r.ExpireDT,
	}

	if !valid.IsNil(r.Category) {
		item.Category, _ = dao.NewItemCategory(int(*r.Category))
	}

	if !valid.IsNil(r.Size) {
		item.Size, _ = dao.NewItemSize(int(*r.Size))
	}

	return item, nil
}

func (u *UpdateItem) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	if !valid.IsNil(u.Category) {
		result["category"] = *u.Category
	}

	if !valid.IsNil(u.Barcode) {
		result["barcode"] = *u.Barcode
	}

	if !valid.IsNil(u.Price) {
		result["price"] = *u.Price
	}

	if !valid.IsNil(u.Cost) {
		result["cost"] = *u.Cost
	}

	if !valid.IsNil(u.Name) {
		result["name"] = *u.Name

		var consonant string
		n := strings.TrimSpace(*u.Name)
		if n != "" {
			han := hangul.ExtractHangul(n)

			for _, h := range han {
				consonant += h.Chosung
			}
		}

		result["consonant"] = consonant
	}

	if !valid.IsNil(u.Description) {
		result["description"] = *u.Description
	}

	if !valid.IsNil(u.ExpireDT) {
		result["expire_dt"] = *u.ExpireDT
	}

	if !valid.IsNil(u.Size) {
		result["size"] = *u.Size
	}

	return result
}
