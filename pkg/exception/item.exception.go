package exception

type GetItemList struct{}

func (e *GetItemList) Error() string {
	return "get item list failed"
}
