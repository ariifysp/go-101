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

	ItemFilter struct {
		Name        string `query:"name" validate:"omitempty,max=64"`
		Description string `query:"description" validate:"omitempty,max=128"`
		Paginate
	}

	Paginate struct {
		Page int64 `query:"page" validate:"required,gte=1"`
		Size int64 `query:"size" validate:"required,min=1,max=20"`
	}
)
