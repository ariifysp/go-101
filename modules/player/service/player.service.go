package service

type PlayerServiceInterface interface{}

type PlayerService struct{}

func NewPlayerService() PlayerServiceInterface {
	return &PlayerService{}
}
