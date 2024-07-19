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
		GetItemByID(itemID uint64) (*itemModel.Item, error)
		CreateItem(item *itemModel.Item) (*itemModel.Item, error)
		UpdateItem(itemID uint64, item *itemModel.ItemUpdate) (uint64, error)
		DeleteItem(itemID uint64) error
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

func (r *ItemRepository) GetItemByID(itemID uint64) (*itemModel.Item, error) {
	itemResult := new(itemModel.Item)

	if err := r.db.Connect().Model(&itemModel.Item{}).First(itemResult, itemID).Error; err != nil {
		return nil, &exception.ItemNotFound{ItemID: itemID}
	}

	return itemResult, nil
}

func (r *ItemRepository) CreateItem(item *itemModel.Item) (*itemModel.Item, error) {
	newItem := new(itemModel.Item)

	if err := r.db.Connect().Model(&itemModel.Item{}).Create(item).Scan(newItem).Error; err != nil {
		return nil, &exception.CreateItem{}
	}

	return newItem, nil
}

func (r *ItemRepository) UpdateItem(itemID uint64, itemUpdate *itemModel.ItemUpdate) (uint64, error) {
	if err := r.db.Connect().Model(&itemModel.Item{}).Where("id = ?", itemID).Updates(itemUpdate).Error; err != nil {
		return 0, &exception.UpdateItem{ItemID: itemID}
	}

	return itemID, nil
}

func (r *ItemRepository) DeleteItem(itemID uint64) error {
	if err := r.db.Connect().Model(&itemModel.Item{}).Where("id = ?", itemID).Update("is_deleted", true).Error; err != nil {
		return &exception.DeleteItem{ItemID: itemID}
	}

	return nil
}
