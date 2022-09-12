package services

import (
	"time"

	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"

	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type VotingService interface {
	CreateVote(vote dto.Vote) (id common.ID_t, err error)
	GetVote(voteID common.ID_t) (*dto.Vote, error)
	GetVotes(pollID common.ID_t) ([]*dto.Vote, error)
	DeleteVote(voteID common.ID_t) (id common.ID_t, err error)
}

type votingService struct {
	repository repository.DAO
}

func NewVotingService(repo repository.DAO) VotingService {
	return &votingService{repository: repo}
}

func (v *votingService) CreateVote(vote dto.Vote) (id common.ID_t, err error) {
	voteModel := &datastruct.Vote{
		PollID:    vote.PollID,
		UserID:    vote.UserID,
		OptionID:  vote.OptionID,
		Timestamp: time.Now().Format(time.RFC850),
	}
	id, err = v.repository.NewVoteQuery().CreateVote(*voteModel)
	return id, nil
}

func (v *votingService) GetVote(voteID common.ID_t) (*dto.Vote, error) {
	voteModel, err := v.repository.NewVoteQuery().GetVote(voteID)
	if err != nil {
		return nil, err
	}
	vote := dto.Vote{
		PollID:    voteModel.PollID,
		Timestamp: voteModel.Timestamp,
	}
	return &vote, nil
}

func (v *votingService) GetVotes(pollID common.ID_t) ([]*dto.Vote, error) {
	votes, _ := v.repository.NewVoteQuery().GetVotes(pollID)
	_ = votes
	return nil, nil
}

func (v *votingService) DeleteVote(voteID common.ID_t) (common.ID_t, error) {
	vID, _ := v.repository.NewVoteQuery().DeleteVote(voteID)
	return vID, nil
}
