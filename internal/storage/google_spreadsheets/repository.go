package spreadsheets

import (
	"fmt"

	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
)

type Storage struct {
}

func NewStorage() (*Storage, error) {
	return &Storage{}, nil
}

func (s *Storage) NewVoteQuery() repository.VoteQuery {
	fmt.Print("New vote query ccalled from spreadsheet\n")
	return &voteQuery{storage: s}
}

func (s *Storage) NewElectionQuery() repository.ElectionQuery {
	return &electionQuery{storage: s}
}
