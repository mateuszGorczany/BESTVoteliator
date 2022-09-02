package database

import (
	"fmt"

	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
)

type Storage struct {
}

func NewStorage() (*Storage, error) {
	return nil, nil
}

func (s *Storage) NewVoteQuery() repository.VoteQuery {
	fmt.Print("New VoteQuery for DB\n")
	return &voteQuery{}
}

func (s *Storage) NewElectionQuery() repository.ElectionQuery {
	return &electionQuery{}
}
