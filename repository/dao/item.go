package dao

import (
	"fmt"
	"strings"
	"time"

	"github.com/LoperLee/golang-hangul-toolkit/hangul"
	"github.com/pkg/errors"
	"hello-cafe/internal/apierror"
	"hello-cafe/model/request"
)

type ItemCategory int

const (
	ItemCategoryBeverage ItemCategory = iota
	ItemCategoryFood
)

func NewItemCategory(category int) (*ItemCategory, error) {
	itemCategory := ItemCategory(category)
	switch itemCategory {
	case ItemCategoryBeverage, ItemCategoryFood:
		return &itemCategory, nil
	default:
		return nil, apierror.ErrNilCategory.SetInternal(fmt.Errorf("category(%d) is invalid", category))
	}
}

type ItemSize int

const (
	ItemSizeSmall ItemSize = iota
	ItemSizeLarge
)

func NewItemSize(size int) (*ItemSize, error) {
	itemSize := ItemSize(size)
	switch itemSize {
	case ItemSizeSmall, ItemSizeLarge:
		return &itemSize, nil
	default:
		return nil, apierror.ErrNilCategory.SetInternal(fmt.Errorf("size(%d) is invalid", size))
	}
}

type Items []Item

type Item struct {
	ItemSeq     int64        `gorm:"Column:item_seq;PRIMARY_KEY"`
	AdminSeq    int64        `gorm:"Column:admin_seq"`
	Category    ItemCategory `gorm:"Column:category"`
	Barcode     string       `gorm:"Column:barcode"`
	Price       int64        `gorm:"Column:price"`
	Cost        int64        `gorm:"Column:cost"`
	Name        string       `gorm:"Column:name"`
	Consonant   string       `gorm:"Column:consonant"`
	Description string       `gorm:"Column:description"`
	ExpireDT    time.Time    `gorm:"Column:expire_dt"`
	Size        ItemSize     `gorm:"Column:size"`
	RegDT       time.Time    `gorm:"Column:reg_dt"`
	ModDT       time.Time    `gorm:"Column:mod_dt"`
}

func NewItem(r request.CreateItem) (*Item, error) {
	if err := r.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to create new item")
	}

	category, _ := NewItemCategory(int(*r.Category))
	size, _ := NewItemSize(int(*r.Size))
	now := time.Now()

	item := &Item{
		AdminSeq:    r.AdminSeq,
		Category:    *category,
		Barcode:     *r.Barcode,
		Price:       *r.Price,
		Cost:        *r.Cost,
		Name:        *r.Name,
		Description: *r.Description,
		ExpireDT:    *r.ExpireDT,
		Size:        *size,
		RegDT:       now,
		ModDT:       now,
	}

	item.SetConsonant()

	return item, nil
}

func (i *Item) SetConsonant() {
	n := strings.TrimSpace(i.Name)

	if n != "" {
		han := hangul.ExtractHangul(n)

		for _, h := range han {
			i.Consonant += h.Chosung
		}
	}
}

func (i *Item) TableName() string {
	return "item"
}
