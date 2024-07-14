package controller

type PlayerControllerInterface interface{}

type PlayerController struct{}

func NewPlayerController() PlayerControllerInterface {
	return &PlayerController{}
}
