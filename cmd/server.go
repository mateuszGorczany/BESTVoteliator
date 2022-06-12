package main

import (
	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"
	"github.com/mateuszGorczany/BESTVoteliator/internal/service"
	database "github.com/mateuszGorczany/BESTVoteliator/internal/storage/postgres_database"
)

type App struct {
	storage       repository.DAO
	votingService service.VotingService
	authService   service.OAuth2AuthenticationService
}

func NewApp(
	storage repository.DAO,
) (*App, error) {
	return &App{
		storage:       storage,
		votingService: service.NewVotingService(storage),
		authService:   service.NewOAuth2AuthenticationService(storage),
	}, nil
}

func main() {
	// storage, _ := spreadsheets.NewStorage()
	storage, _ := database.NewStorage()

	app, _ := NewApp(storage)
	// votingService2.CreateVote(dto.Vote{})
	app.votingService.CreateVote(dto.Vote{})
	app.votingService.GetVote(1234)
	_ = app
}
