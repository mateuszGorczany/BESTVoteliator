package database

import (
	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type electionQuery struct {
	storage *Storage
}

func (e *electionQuery) CreateElection(election datastruct.Election) (id common.ID_t, err error) {
	return common.ID_t(12), nil
}
func (e *electionQuery) GetElection(id common.ID_t) (*datastruct.Election, error) {
	return nil, nil
}
func (e *electionQuery) GetElections() ([]*datastruct.Election, error) {
	return nil, nil
}
func (e *electionQuery) UpdateElection() *datastruct.Election {
	return nil
}
func (e *electionQuery) DeleteElection(id common.ID_t) common.ID_t {
	return common.ID_t(0)
}
