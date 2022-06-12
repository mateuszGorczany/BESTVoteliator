package database

import (
	"fmt"

	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
)

type voteQuery struct{}

func (v *voteQuery) CreateVote() (id int64, err error) {
	fmt.Print("HelloDB\n")
	var i int = 0
	return int64(i), nil
}

func (v *voteQuery) GetVote() (*datastruct.Vote, error) {
	return &datastruct.Vote{}, nil
}

func (v *voteQuery) GetVotes() ([]*datastruct.Vote, error) {

	return []*datastruct.Vote{{}, {}}, nil
}

func (v *voteQuery) UpdateVote() *datastruct.Vote {

	return &datastruct.Vote{}
}

func (v *voteQuery) DeleteVote() (id int64) {
	return 0
}
