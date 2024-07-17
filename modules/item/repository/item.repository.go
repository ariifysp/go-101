package repository

import (
	itemModel "github/ariifysp/go-101/modules/item/model"
	"github/ariifysp/go-101/pkg/database"
	"github/ariifysp/go-101/pkg/exception"
)

type (
	ItemRepositoryInterface interface {
		GetItemList(itemFilter *itemModel.ItemFilter) ([]*itemModel.Item, error)
		CountItems(itemFilter *itemModel.ItemFilter) (int64, error)
	}

	ItemRepository struct {
		db database.DatabaseInterface
	}
)

func NewItemRepository(db database.DatabaseInterface) ItemRepositoryInterface {
	return &ItemRepository{db}
}

func (r *ItemRepository) GetItemList(itemFilter *itemModel.ItemFilter) ([]*itemModel.Item, error) {
	items := make([]*itemModel.Item, 0)

	query := r.db.Connect().Model(&itemModel.Item{}).Where("is_deleted = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	limit := int(itemFilter.Size)
	offset := int((itemFilter.Page - 1) * itemFilter.Size)

	if err := query.Limit(limit).Offset(offset).Find(&items).Order("id asc").Error; err != nil {
		return nil, &exception.GetItemList{}
	}

	return items, nil
}

func (r *ItemRepository) CountItems(itemFilter *itemModel.ItemFilter) (int64, error) {
	var count int64

	query := r.db.Connect().Model(&itemModel.Item{}).Where("is_deleted = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return -1, &exception.CountItems{}
	}

	return count, nil
}
