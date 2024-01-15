package repository

import (
	"strings"

	"github.com/pkg/errors"
	"hello-cafe/internal/apierror"
	"hello-cafe/internal/db"
	"hello-cafe/internal/valid"
	"hello-cafe/model/request"
	"hello-cafe/repository/dao"
)

type ItemRepository interface {
	Create(item request.CreateItem) error
	Update(item request.UpdateItem) error
	Delete(itemSeq int64) error
	Find(adminSeq, lastItemSeq int64, limit int) (dao.Items, error)
	Get(itemSeq int64) (*dao.Item, error)
	GetByBarcode(barcode string) (*dao.Item, error)
	Search(adminSeq int64, text string) (dao.Items, error)
}

type itemRepository struct{}

func NewItemRepository() ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) Create(item request.CreateItem) error {
	newItem, err := dao.NewItem(item)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := db.Conn().Create(&newItem).Error; err != nil {
		return errors.Wrap(err, "failed to create item")
	}

	return nil
}

func (r *itemRepository) Update(item request.UpdateItem) error {
	updateItem, err := NewUpdateItem(item)
	if err != nil {
		return errors.WithStack(err)
	}

	i, err := r.Get(item.ItemSeq)
	if err != nil {
		return errors.Wrap(err, "failed to get item")
	}

	if valid.IsNil(i) {
		return apierror.ErrNotExistItem
	}

	if err := db.Conn().Model(&i).Updates(updateItem.ToMap()).Error; err != nil {
		return errors.Wrap(err, "failed to update item info")
	}

	return nil
}

func (r *itemRepository) Delete(itemSeq int64) error {
	item, err := r.Get(itemSeq)
	if err != nil {
		return errors.Wrap(err, "failed to get item")
	}

	if valid.IsNil(item) {
		return apierror.ErrNotExistItem
	}

	return errors.WithStack(db.Conn().Delete(&item).Error)
}

func (r *itemRepository) Find(adminSeq int64, lastItemSeq int64, limit int) (dao.Items, error) {
	if adminSeq < 0 {
		return nil, apierror.ErrInvalidAdmin
	}

	if limit <= 0 {
		limit = 10
	}

	tx := db.Conn().
		Table("item").
		Select("*").
		Limit(limit).
		Order("item_seq DESC")

	if lastItemSeq > 0 {
		tx = tx.Where("item_seq < ?", lastItemSeq)
	}

	items := make(dao.Items, 0)
	if err := tx.Find(&items).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find items")
	}

	return items, nil
}

func (r *itemRepository) Get(itemSeq int64) (*dao.Item, error) {
	if itemSeq < 0 {
		return nil, apierror.ErrInvalidItem
	}

	var item dao.Item
	if err := db.Conn().Take(&item, itemSeq).Error; err != nil {
		return nil, errors.Wrap(err, "failed to take item info")
	}

	return &item, nil
}

func (r *itemRepository) GetByBarcode(barcode string) (*dao.Item, error) {
	var item dao.Item
	if err := db.Conn().Where("barcode = ?", barcode).Take(&item).Error; err != nil {
		return nil, errors.Wrap(err, "failed to take item info")
	}

	return &item, nil
}

func (r *itemRepository) Search(adminSeq int64, text string) (dao.Items, error) {
	switch {
	case adminSeq < 0:
		return nil, apierror.ErrInvalidAdmin
	case len(strings.TrimSpace(text)) == 0:
		return nil, errors.New("text is empty")
	}

	consonant := "%" + text + "%"
	tx := db.Conn().
		Table("item").
		Select("*").
		Where("admin_seq = ?", adminSeq).
		Where("(name LIKE ? OR consonant LIKE ?)", consonant, consonant).
		Order("item_seq DESC")

	items := make(dao.Items, 0)
	if err := tx.Find(&items).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search items")
	}

	return items, nil
}
