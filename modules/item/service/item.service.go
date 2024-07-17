package service

import (
	itemModel "github/ariifysp/go-101/modules/item/model"
	itemRepository "github/ariifysp/go-101/modules/item/repository"
	"math"
)

type (
	ItemServiceInterface interface {
		ItemListService(itemFilter *itemModel.ItemFilter) (*itemModel.ItemResult, error)
	}

	ItemService struct {
		itemRepository itemRepository.ItemRepositoryInterface
	}
)

func NewItemService(itemRepository itemRepository.ItemRepositoryInterface) ItemServiceInterface {
	return &ItemService{itemRepository}
}

func (s *ItemService) ItemListService(itemFilter *itemModel.ItemFilter) (*itemModel.ItemResult, error) {
	total, err := s.itemRepository.CountItems(itemFilter)
	if err != nil {
		return nil, err
	}

	itemList, err := s.itemRepository.GetItemList(itemFilter)
	if err != nil {
		return nil, err
	}

	page := itemFilter.Page
	size := itemFilter.Size
	totalPage := int64(math.Ceil(float64(total) / float64(size)))

	result := s.itemsResultResponse(itemList, page, totalPage)

	return result, nil
}

func (s *ItemService) itemsResultResponse(itemList []*itemModel.Item, page, totalPage int64) *itemModel.ItemResult {
	items := make([]*itemModel.ItemModel, 0)
	for _, item := range itemList {
		items = append(items, item.ToItemModel())
	}

	return &itemModel.ItemResult{
		Items: items,
		Paginate: itemModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
