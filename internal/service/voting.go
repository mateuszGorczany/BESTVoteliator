package service

import (
	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"

	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type VotingService interface {
	CreateVote(vote dto.Vote) (id *int, err error)
	GetVote(id int) (*datastruct.Vote, error)
	GetVotes() ([]*datastruct.Vote, error)
	CreateElection(election dto.Election) (id *int, err error)
}

type votingService struct {
	repository repository.DAO
}

func NewVotingService(repo repository.DAO) VotingService {
	return &votingService{repository: repo}
}

func (v *votingService) CreateVote(vote dto.Vote) (id *int, err error) {
	v.repository.NewVoteQuery().CreateVote()
	return nil, common.ErrorNotImplementedYet
}

func (v *votingService) GetVote(id int) (*datastruct.Vote, error) {
	return nil, common.ErrorNotImplementedYet
}

func (v *votingService) GetVotes() ([]*datastruct.Vote, error) {
	return nil, common.ErrorNotImplementedYet
}

// creates new datastruct.Elecction
func (v *votingService) CreateElection(election dto.Election) (id *int, err error) {
	return nil, common.ErrorNotImplementedYet
}
