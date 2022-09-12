package repository

import (
	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type VoteQuery interface {
	CreateVote(datastruct.Vote) (id common.ID_t, err error)
	GetVote(voteID common.ID_t) (*datastruct.Vote, error)
	GetVotes(pollID common.ID_t) ([]*datastruct.Vote, error)
	UpdateVote() *datastruct.Vote
	DeleteVote(voteID common.ID_t) (common.ID_t, error)
}
