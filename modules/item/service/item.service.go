package service

import (
	itemModel "github/ariifysp/go-101/modules/item/model"
	itemRepository "github/ariifysp/go-101/modules/item/repository"
)

type (
	ItemServiceInterface interface {
		ItemListService(itemFilter *itemModel.ItemFilter) ([]*itemModel.Item, error)
	}

	ItemService struct {
		itemRepository itemRepository.ItemRepositoryInterface
	}
)

func NewItemService(itemRepository itemRepository.ItemRepositoryInterface) ItemServiceInterface {
	return &ItemService{itemRepository}
}

func (s *ItemService) ItemListService(itemFilter *itemModel.ItemFilter) ([]*itemModel.Item, error) {
	itemList, err := s.itemRepository.GetItemList(itemFilter)
	if err != nil {
		return nil, err
	}

	return itemList, nil
}
