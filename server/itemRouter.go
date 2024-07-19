package server

import (
	itemController "github/ariifysp/go-101/modules/item/controller"
	itemRepository "github/ariifysp/go-101/modules/item/repository"
	itemServie "github/ariifysp/go-101/modules/item/service"
)

func (s *Server) ItemRouter() {
	routerItem := s.App.Group("/v1/item")

	itemRepository := itemRepository.NewItemRepository(s.DB)
	itemService := itemServie.NewItemService(itemRepository)
	itemController := itemController.NewItemController(itemService)

	routerItem.Get("", itemController.ItemList)
	routerItem.Post("", itemController.CreateItem)
	routerItem.Patch("/:itemID", itemController.UpdateItem)
	routerItem.Delete("/:itemID", itemController.DeleteItem)
}
