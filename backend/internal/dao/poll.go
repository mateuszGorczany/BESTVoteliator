package repository

import (
	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type PollQuery interface {
	CreatePoll(datastruct.Poll) (id common.ID_t, err error)
	GetPoll(common.ID_t) (*datastruct.Poll, error)
	GetPolls() ([]*datastruct.Poll, error)
	UpdatePoll() *datastruct.Poll
	DeletePoll(common.ID_t) (common.ID_t, error)
}
