package service

import (
	itemModel "github/ariifysp/go-101/modules/item/model"
	itemRepository "github/ariifysp/go-101/modules/item/repository"
	"math"
)

type (
	ItemServiceInterface interface {
		ItemList(itemFilter *itemModel.ItemFilter) (*itemModel.ItemResult, error)
		CreateItem(itemCreate *itemModel.ItemCreate) (*itemModel.Item, error)
		UpdateItem(itemID uint64, itemUpdate *itemModel.ItemUpdate) (*itemModel.Item, error)
		DeleteItem(itemID uint64) error
	}

	ItemService struct {
		itemRepository itemRepository.ItemRepositoryInterface
	}
)

func NewItemService(itemRepository itemRepository.ItemRepositoryInterface) ItemServiceInterface {
	return &ItemService{itemRepository}
}

func (s *ItemService) ItemList(itemFilter *itemModel.ItemFilter) (*itemModel.ItemResult, error) {
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

func (s *ItemService) CreateItem(itemCreate *itemModel.ItemCreate) (*itemModel.Item, error) {
	item := &itemModel.Item{
		Name:        itemCreate.Name,
		Description: itemCreate.Description,
		Picture:     itemCreate.Picture,
		Price:       itemCreate.Price,
	}

	newItem, err := s.itemRepository.CreateItem(item)
	if err != nil {
		return nil, err
	}

	return newItem, nil
}

func (s *ItemService) UpdateItem(itemID uint64, itemUpdate *itemModel.ItemUpdate) (*itemModel.Item, error) {
	_, err := s.itemRepository.UpdateItem(itemID, itemUpdate)
	if err != nil {
		return nil, err
	}

	item, err := s.itemRepository.GetItemByID(itemID)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemService) DeleteItem(itemID uint64) error {
	return s.itemRepository.DeleteItem(itemID)
}
