package request

import (
	"time"

	"github.com/pkg/errors"
	"hello-cafe/internal/apierror"
	"hello-cafe/internal/valid"
)

type ItemCategory int

const (
	ItemCategoryBeverage ItemCategory = iota
	ItemCategoryFood
)

func (c ItemCategory) Validate() error {
	switch c {
	case ItemCategoryBeverage, ItemCategoryFood:
		return nil
	default:
		return apierror.ErrInvalidCategory
	}
}

type ItemSize int

const (
	ItemSizeSmall ItemSize = iota
	ItemSizeLarge
)

func (s ItemSize) Validate() error {
	switch s {
	case ItemSizeSmall, ItemSizeLarge:
		return nil
	default:
		return apierror.ErrInvalidSize
	}
}

type CreateItem struct {
	AdminSeq    int64         `json:"admin_seq"`
	Category    *ItemCategory `json:"category"`
	Barcode     *string       `json:"barcode"`
	Price       *int64        `json:"price"`
	Cost        *int64        `json:"cost"`
	Name        *string       `json:"name"`
	Description *string       `json:"description"`
	ExpireDT    *time.Time    `json:"expire_dt"`
	Size        *ItemSize     `json:"size"`
}

func (i *CreateItem) Validate() error {
	switch {
	case i.AdminSeq < 0:
		return apierror.ErrInvalidAdmin
	case valid.IsNil(i.Category):
		return apierror.ErrNilCategory
	case valid.IsNil(i.Barcode):
		return apierror.ErrNilBarcode
	case valid.IsNil(i.Price):
		return apierror.ErrNilPrice
	case valid.IsNil(i.Cost):
		return apierror.ErrNilCost
	case valid.IsNil(i.Name):
		return apierror.ErrNilName
	case valid.IsNil(i.Description):
		return apierror.ErrNilDescription
	case valid.IsNil(i.ExpireDT):
		return apierror.ErrNilExpireDT
	case valid.IsNil(i.Size):
		return apierror.ErrNilSize
	}

	if err := i.Category.Validate(); err != nil {
		return errors.Wrapf(err, "category(%d) is invalid", i.Category)
	}

	if err := i.Size.Validate(); err != nil {
		return errors.Wrapf(err, "size(%d) is invalid", i.Size)
	}

	return nil
}

type UpdateItem struct {
	ItemSeq     int64         `uri:"item_seq"`
	Category    *ItemCategory `json:"category"`
	Barcode     *string       `json:"barcode"`
	Price       *int64        `json:"price"`
	Cost        *int64        `json:"cost"`
	Name        *string       `json:"name"`
	Description *string       `json:"description"`
	ExpireDT    *time.Time    `json:"expire_dt"`
	Size        *ItemSize     `json:"size"`
}

func (i *UpdateItem) Validate() error {
	if i.ItemSeq < 0 {
		return apierror.ErrInvalidItem
	}

	if !valid.IsNil(i.Category) {
		if err := i.Category.Validate(); err != nil {
			return errors.Wrapf(err, "category(%d) is invalid", i.Category)
		}
	}

	if !valid.IsNil(i.Size) {
		if err := i.Size.Validate(); err != nil {
			return errors.Wrapf(err, "size(%d) is invalid", i.Size)
		}
	}

	return nil
}

type DeleteItem struct {
	ItemSeq int64 `uri:"item_seq"`
}

func (i *DeleteItem) Validate() error {
	if i.ItemSeq <= 0 {
		return apierror.ErrInvalidItem
	}

	return nil
}

type GetItem struct {
	ItemSeq int64 `uri:"item_seq"`
}

func (i *GetItem) Validate() error {
	if i.ItemSeq <= 0 {
		return apierror.ErrInvalidItem
	}

	return nil
}
