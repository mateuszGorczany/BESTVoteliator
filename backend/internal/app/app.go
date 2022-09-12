package application

import (
	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	"github.com/mateuszGorczany/BESTVoteliator/internal/services"
	database "github.com/mateuszGorczany/BESTVoteliator/internal/storage/firebase_database"
)

type MicroserviceApp struct {
	storage        repository.DAO
	VotingService  services.VotingService
	PollingService services.PollingService
	AuthService    services.OAuth2AuthenticationService
}

func NewMicroserviceApp() (*MicroserviceApp, error) {
	storage, _ := database.NewStorage()
	return &MicroserviceApp{
		storage:        storage,
		VotingService:  services.NewVotingService(storage),
		PollingService: services.NewPollingService(storage),
		AuthService:    services.NewOAuth2AuthenticationService(storage),
	}, nil
}
