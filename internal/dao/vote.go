package repository

import "github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"

type VoteQuery interface {
	CreateVote() (id int64, err error)
	GetVote() (*datastruct.Vote, error)
	GetVotes() ([]*datastruct.Vote, error)
	UpdateVote() *datastruct.Vote
	DeleteVote() (id int64)
}
