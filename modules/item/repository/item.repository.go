package repository

import (
	itemModel "github/ariifysp/go-101/modules/item/model"
	"github/ariifysp/go-101/pkg/database"
	"github/ariifysp/go-101/pkg/exception"
)

type (
	ItemRepositoryInterface interface {
		GetItemList(itemFilter *itemModel.ItemFilter) ([]*itemModel.Item, error)
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
	if err := query.Find(&items).Order("id asc").Error; err != nil {
		return nil, &exception.GetItemList{}
	}

	return items, nil
}
