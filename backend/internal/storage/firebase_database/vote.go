package database

import (
	"context"
	"fmt"

	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
	"google.golang.org/api/iterator"
)

type voteQuery struct{}

func (v *voteQuery) CreateVote(vote datastruct.Vote) (id common.ID_t, err error) {
	fmt.Printf("vote: %v\n", vote)
	doc, _, err := VoteQueryBuilder().Add(context.Background(), vote)
	if err != nil {
		return common.ID_t(doc.ID), err
	}
	return common.ID_t(doc.ID), nil
}

func (v *voteQuery) GetVote(id common.ID_t) (*datastruct.Vote, error) {
	doc, err := VoteQueryBuilder().Doc(string(id)).Get(context.Background())
	if err != nil {
		return nil, err
	}
	vote := &datastruct.Vote{}
	doc.DataTo(vote)
	return vote, nil
}

func (v *voteQuery) GetVotes(pollID common.ID_t) ([]*datastruct.Vote, error) {
	documentIterator := VoteQueryBuilder().Documents(context.Background())
	votes := make([]*datastruct.Vote, 0)
	i := 0
	for {
		doc, err := documentIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		vote := &datastruct.Vote{}
		doc.DataTo(vote)
		votes = append(votes, vote)
		i++
	}
	return votes, nil
}

func (v *voteQuery) UpdateVote() *datastruct.Vote {

	return &datastruct.Vote{}
}

func (v *voteQuery) DeleteVote(id common.ID_t) (common.ID_t, error) {
	return "", common.ErrorNotImplemented
}
