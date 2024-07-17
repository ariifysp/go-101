package exception

type GetItemList struct{}
type CountItems struct{}

func (e *GetItemList) Error() string {
	return "get item list failed"
}

func (e *CountItems) Error() string {
	return "count items failed"
}
