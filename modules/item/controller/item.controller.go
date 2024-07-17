package controller

import (
	itemModel "github/ariifysp/go-101/modules/item/model"
	itemService "github/ariifysp/go-101/modules/item/service"
	"github/ariifysp/go-101/pkg/custom"

	"github.com/gofiber/fiber/v3"
)

type (
	ItemControllerInterface interface {
		ItemList(ctx fiber.Ctx) error
	}

	ItemController struct {
		itemService itemService.ItemServiceInterface
	}
)

func NewItemController(itemService itemService.ItemServiceInterface) ItemControllerInterface {
	return &ItemController{itemService}
}

func (c *ItemController) ItemList(ctx fiber.Ctx) error {
	itemFilter := new(itemModel.ItemFilter)

	customFiberRequest := custom.NewCustomFiberRequest(ctx)
	if err := customFiberRequest.BindQuery(itemFilter); err != nil {
		return custom.ErrorResponse(ctx, fiber.ErrBadRequest, err)
	}

	result, err := c.itemService.ItemListService(itemFilter)
	if err != nil {
		return custom.ErrorResponse(ctx, fiber.ErrInternalServerError, err)
	}

	return custom.DataResponse(ctx, fiber.StatusOK, result)
}
