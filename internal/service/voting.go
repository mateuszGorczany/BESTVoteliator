package service

import (
	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"

	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type VotingService interface {
	CreateVote(vote dto.Vote) (id common.ID_t, err error)
	GetVote(id common.ID_t) (*dto.Vote, error)
	GetVotes() ([]*dto.Vote, error)
	CreateElection(election dto.Election) (id common.ID_t, err error)
}

type votingService struct {
	repository repository.DAO
}

func NewVotingService(repo repository.DAO) VotingService {
	return &votingService{repository: repo}
}

func (v *votingService) CreateVote(vote dto.Vote) (id common.ID_t, err error) {
	id, err = v.repository.NewVoteQuery().CreateVote(datastruct.Vote{})
	return id, nil
}

func (v *votingService) GetVote(id common.ID_t) (*dto.Vote, error) {
	// vote, _ := v.repository.NewVoteQuery().GetVote()

	// return &dto.Vote{
	// 	ID:     vote.ID,
	// 	UserID: vote.UserID,
	// 	// time.April.String(),
	// }, common.ErrorNotImplementedYet
	return &dto.Vote{ID: int64(1337),
		UserID: int64(2134)}, nil
}

func (v *votingService) GetVotes() ([]*dto.Vote, error) {
	votes, _ := v.GetVotes()
	return votes, nil
}

// creates new datastruct.Elecction
func (v *votingService) CreateElection(election dto.Election) (id common.ID_t, err error) {
	id, err = v.repository.NewElectionQuery().CreateElection(datastruct.Election{})
	return id, nil
}
