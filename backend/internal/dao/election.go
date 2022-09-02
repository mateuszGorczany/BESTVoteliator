package repository

import (
	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type ElectionQuery interface {
	CreateElection(datastruct.Election) (id common.ID_t, err error)
	GetElection(common.ID_t) (*datastruct.Election, error)
	GetElections() ([]*datastruct.Election, error)
	UpdateElection() *datastruct.Election
	DeleteElection(common.ID_t) (id common.ID_t)
}
