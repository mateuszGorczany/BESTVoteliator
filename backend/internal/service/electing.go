package service

import (
	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"

	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type ElectingService interface {
	CreateElection(election dto.Election) (common.ID_t, error)
	GetElection(id common.ID_t) (*dto.Election, error)
	GetElections() ([]*dto.Election, error)
	DeleteElection(id common.ID_t) (common.ID_t, error)
}

type electingService struct {
	repository repository.DAO
}

func NewElectingService(repo repository.DAO) ElectingService {
	return &electingService{repository: repo}
}

func (e *electingService) CreateElection(vote dto.Election) (id common.ID_t, err error) {
	id, err = e.repository.NewVoteQuery().CreateVote(datastruct.Vote{})
	_ = err
	return id, common.ErrorNotImplementedYet
}

func (e *electingService) GetElection(id common.ID_t) (*dto.Election, error) {
	// vote, _ := v.repository.NewVoteQuery().GetVote()

	// return &dto.Vote{
	// 	ID:     vote.ID,
	// 	UserID: vote.UserID,
	// 	// time.April.String(),
	// }, common.ErrorNotImplementedYet
	return &dto.Election{}, common.ErrorNotImplementedYet
}

func (e *electingService) GetElections() ([]*dto.Election, error) {
	elections, err := e.GetElections()
	_ = err
	return elections, nil
}

func (e *electingService) DeleteElection(id common.ID_t) (common.ID_t, error) {
	return common.ID_t(0), nil
}
