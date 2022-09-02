package spreadsheets

import (
	"fmt"

	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type voteQuery struct {
	storage *Storage
}

func (v *voteQuery) CreateVote(vote datastruct.Vote) (id common.ID_t, err error) {
	fmt.Print("Create vote query ccalled from voteQuery\n")
	var i int = 0
	return common.ID_t(i), nil
}

func (v *voteQuery) GetVote(id common.ID_t) (*datastruct.Vote, error) {
	fmt.Print("Create vote query ccalled from voteQuery\n")
	return &datastruct.Vote{
		ID:     int64(1234),
		UserID: int64(4566),
	}, nil
}

func (v *voteQuery) GetVotes() ([]*datastruct.Vote, error) {

	return []*datastruct.Vote{{}, {}}, nil
}

func (v *voteQuery) UpdateVote() *datastruct.Vote {

	return &datastruct.Vote{}
}

func (v *voteQuery) DeleteVote(id common.ID_t) common.ID_t {
	return 0
}
