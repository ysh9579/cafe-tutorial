package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"hello-cafe/internal/apierror"
	"hello-cafe/internal/valid"
	"hello-cafe/model"
	"hello-cafe/model/request"
	"hello-cafe/repository"
	"hello-cafe/repository/dao"
)

type ItemService interface {
	Create(item request.CreateItem) error
	Update(item request.UpdateItem) error
	Delete(itemSeq int64) error
	Find(adminSeq, lastItemSeq int64, limit int) (model.Items, error)
	Get(itemSeq int64) (*model.Item, error)
	Search(adminSeq int64, text string) (model.Items, error)
	CheckDuplicated(barcode string) (bool, error)
}

type itemService struct {
	repo repository.Repository
}

func NewItemService(repo repository.Repository) (ItemService, error) {
	if valid.IsNil(repo) {
		return nil, errors.New("repository is nil")
	}

	return &itemService{repo: repo}, nil
}

func (s *itemService) Create(item request.CreateItem) error {
	if err := item.Validate(); err != nil {
		return errors.WithStack(err)
	}

	if _, err := s.repo.Admin().Get(item.AdminSeq); err != nil {
		return apierror.ErrInvalidAdmin
	}

	isDuplicated, err := s.CheckDuplicated(*item.Barcode)
	if err != nil {
		return errors.WithStack(err)
	}

	if isDuplicated {
		return apierror.ErrDuplicatedItem
	}

	if err := s.repo.Item().Create(item); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *itemService) CheckDuplicated(barcode string) (bool, error) {
	item, err := s.repo.Item().GetByBarcode(barcode)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.WithStack(err)
	}

	if !valid.IsNil(item) && item.ItemSeq > 0 {
		return true, nil
	}

	return false, nil
}

func (s *itemService) Update(item request.UpdateItem) error {
	if err := item.Validate(); err != nil {
		return errors.WithStack(err)
	}

	_, err := s.repo.Item().Get(item.ItemSeq)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithStack(err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apierror.ErrNotExistItem
	}

	if err := s.repo.Item().Update(item); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *itemService) Delete(itemSeq int64) error {
	if itemSeq <= 0 {
		return apierror.ErrInvalidItem
	}

	_, err := s.repo.Item().Get(itemSeq)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithStack(err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apierror.ErrNotExistItem
	}

	if err := s.repo.Item().Delete(itemSeq); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *itemService) Find(adminSeq, lastItemSeq int64, limit int) (model.Items, error) {
	if adminSeq <= 0 {
		return nil, apierror.ErrInvalidAdmin
	}

	if _, err := s.repo.Admin().Get(adminSeq); err != nil {
		return nil, apierror.ErrInvalidAdmin
	}

	daoItems, err := s.repo.Item().Find(adminSeq, lastItemSeq, limit)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "failed to find item list")
	}

	result := make(model.Items, 0)
	for _, item := range daoItems {
		result = append(result, s.getItemFromDAO(item))
	}

	return result, nil
}

func (s *itemService) Get(itemSeq int64) (*model.Item, error) {
	if itemSeq <= 0 {
		return nil, apierror.ErrInvalidItem
	}

	item, err := s.repo.Item().Get(itemSeq)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "failed to get item")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apierror.ErrNotExistItem
	}

	result := s.getItemFromDAO(*item)

	return &result, nil
}

func (s *itemService) getItemFromDAO(item dao.Item) model.Item {
	return model.Item{
		ItemSeq:     item.ItemSeq,
		AdminSeq:    item.AdminSeq,
		Category:    int(item.Category),
		Barcode:     item.Barcode,
		Price:       item.Price,
		Cost:        item.Cost,
		Name:        item.Name,
		Description: item.Description,
		ExpireDT:    item.ExpireDT,
		Size:        int(item.Size),
		RegDT:       item.RegDT,
		ModDT:       item.ModDT,
	}
}

func (s *itemService) Search(adminSeq int64, text string) (model.Items, error) {
	switch {
	case adminSeq <= 0:
		return nil, apierror.ErrInvalidAdmin
	case text == "":
		return nil, apierror.ErrNilSearchText
	}

	if _, err := s.repo.Admin().Get(adminSeq); err != nil {
		return nil, apierror.ErrInvalidAdmin
	}

	daoItems, err := s.repo.Item().Search(adminSeq, text)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search")
	}

	result := make(model.Items, 0)
	for _, item := range daoItems {
		result = append(result, s.getItemFromDAO(item))
	}

	return result, nil
}
