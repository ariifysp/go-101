package repository

type PlayerRepositoryInterface interface{}

type PlayerRepository struct{}

func NewPlayerRepository() PlayerRepositoryInterface {
	return &PlayerRepository{}
}
