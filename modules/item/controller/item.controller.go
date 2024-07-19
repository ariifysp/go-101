package controller

import (
	"fmt"
	itemModel "github/ariifysp/go-101/modules/item/model"
	itemService "github/ariifysp/go-101/modules/item/service"
	"github/ariifysp/go-101/pkg/custom"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type (
	ItemControllerInterface interface {
		ItemList(ctx fiber.Ctx) error
		CreateItem(ctx fiber.Ctx) error
		UpdateItem(ctx fiber.Ctx) error
		DeleteItem(ctx fiber.Ctx) error
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

	result, err := c.itemService.ItemList(itemFilter)
	if err != nil {
		return custom.ErrorResponse(ctx, fiber.ErrInternalServerError, err)
	}

	return custom.DataResponse(ctx, fiber.StatusOK, result)
}

func (c *ItemController) CreateItem(ctx fiber.Ctx) error {
	itemCreate := new(itemModel.ItemCreate)

	customFiberRequest := custom.NewCustomFiberRequest(ctx)
	if err := customFiberRequest.BindBody(itemCreate); err != nil {
		return custom.ErrorResponse(ctx, fiber.ErrBadRequest, err)
	}

	result, err := c.itemService.CreateItem(itemCreate)
	if err != nil {
		return custom.ErrorResponse(ctx, fiber.ErrInternalServerError, err)
	}

	return custom.DataResponse(ctx, fiber.StatusCreated, result)
}

func (c *ItemController) UpdateItem(ctx fiber.Ctx) error {
	itemID, err := c.getItemID(ctx)
	if err != nil {
		return custom.ErrorResponse(ctx, fiber.ErrBadRequest, err)
	}

	itemUpdate := new(itemModel.ItemUpdate)

	customFiberRequest := custom.NewCustomFiberRequest(ctx)
	if err := customFiberRequest.BindBody(itemUpdate); err != nil {
		return custom.ErrorResponse(ctx, fiber.ErrBadRequest, err)
	}

	result, err := c.itemService.UpdateItem(itemID, itemUpdate)
	if err != nil {
		return custom.ErrorResponse(ctx, fiber.ErrInternalServerError, err)
	}

	return custom.DataResponse(ctx, fiber.StatusCreated, result)
}

func (c *ItemController) DeleteItem(ctx fiber.Ctx) error {
	itemID, err := c.getItemID(ctx)
	if err != nil {
		return custom.ErrorResponse(ctx, fiber.ErrBadRequest, err)
	}

	if err := c.itemService.DeleteItem(itemID); err != nil {
		return custom.ErrorResponse(ctx, fiber.ErrInternalServerError, err)
	}

	message := fmt.Sprintf("delete item id: %d success", itemID)
	return custom.SuccessResponse(ctx, fiber.StatusNoContent, message)
}

func (c *ItemController) getItemID(ctx fiber.Ctx) (uint64, error) {
	itemID := ctx.Params("itemID")
	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, err
	}

	return itemIDUint64, nil
}
