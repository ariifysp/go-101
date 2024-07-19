package model

import "time"

type (
	Item struct {
		ID          uint64
		AdminID     *string
		Name        string
		Description string
		Picture     string
		Price       uint
		IsDeleted   bool
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	ItemModel struct {
		ID          uint64 `json:"itemID"`
		Name        string `json:"itemName"`
		Description string `json:"description"`
		Picture     string `json:"picture"`
		Price       uint   `json:"price"`
	}

	ItemCreate struct {
		AdminID     string
		Name        string `json:"name" validate:"required,max=64"`
		Description string `json:"description" validate:"required,max=128"`
		Picture     string `json:"picture" validate:"required"`
		Price       uint   `json:"price" validate:"required"`
	}

	ItemUpdate struct {
		AdminID     string
		Name        string `json:"name" validate:"omitempty,max=64"`
		Description string `json:"description" validate:"omitempty,max=128"`
		Picture     string `json:"picture" validate:"omitempty"`
		Price       uint   `json:"price" validate:"omitempty"`
	}

	ItemFilter struct {
		Name        string `query:"name" validate:"omitempty,max=64"`
		Description string `query:"description" validate:"omitempty,max=128"`
		Paginate
	}

	Paginate struct {
		Page int64 `query:"page" validate:"required,gte=1"`
		Size int64 `query:"size" validate:"required,min=1,max=20"`
	}

	ItemResult struct {
		Items    []*ItemModel   `json:"items"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct {
		Page      int64 `json:"page"`
		TotalPage int64 `json:"totalPage"`
	}
)

func (i *Item) ToItemModel() *ItemModel {
	return &ItemModel{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Picture:     i.Picture,
		Price:       i.Price,
	}
}
