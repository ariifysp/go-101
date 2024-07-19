package exception

import "fmt"

type GetItemList struct{}
type CountItems struct{}
type CreateItem struct{}
type ItemNotFound struct {
	ItemID uint64
}
type UpdateItem struct {
	ItemID uint64
}
type DeleteItem struct {
	ItemID uint64
}

func (e *GetItemList) Error() string {
	return "get item list failed"
}

func (e *CountItems) Error() string {
	return "count items failed"
}

func (e *CreateItem) Error() string {
	return "create item failed"
}

func (e *ItemNotFound) Error() string {
	return fmt.Sprintf("itemID: %d was not found", e.ItemID)
}

func (e *UpdateItem) Error() string {
	return fmt.Sprintf("update item id: %d failed", e.ItemID)
}

func (e *DeleteItem) Error() string {
	return fmt.Sprintf("delete item id: %d failed", e.ItemID)
}
