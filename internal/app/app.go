package application

import (
	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	"github.com/mateuszGorczany/BESTVoteliator/internal/service"
	spreadsheets "github.com/mateuszGorczany/BESTVoteliator/internal/storage/google_spreadsheets"
	database "github.com/mateuszGorczany/BESTVoteliator/internal/storage/postgres_database"
)

type MicroserviceApp struct {
	storage         repository.DAO
	VotingService   service.VotingService
	ElectingService service.ElectingService
	AuthService     service.OAuth2AuthenticationService
}

func NewMicroserviceApp() (*MicroserviceApp, error) {
	storage0, _ := spreadsheets.NewStorage()
	_ = storage0
	storage, _ := database.NewStorage()
	return &MicroserviceApp{
		storage:         storage,
		VotingService:   service.NewVotingService(storage),
		ElectingService: service.NewElectingService(storage),
		AuthService:     service.NewOAuth2AuthenticationService(storage),
	}, nil
}
