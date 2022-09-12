package database

import (
	"context"

	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
	"google.golang.org/api/iterator"
)

type electionQuery struct{}

func (e *electionQuery) CreatePoll(election datastruct.Poll) (id common.ID_t, err error) {
	doc, _, err := PollQueryBuiler().Add(context.Background(), election)
	if err != nil {
		return "", err
	}
	return common.ID_t(doc.ID), nil
}
func (e *electionQuery) GetPoll(id common.ID_t) (*datastruct.Poll, error) {
	doc, err := PollQueryBuiler().Doc(string(id)).Get(context.Background())
	if err != nil {
		return nil, err
	}
	poll := &datastruct.Poll{}
	doc.DataTo(poll)
	return poll, nil
}

func (e *electionQuery) GetPolls() ([]*datastruct.Poll, error) {
	documentIterator := PollQueryBuiler().Documents(context.Background())
	polls := make([]*datastruct.Poll, 0)
	i := 0
	for {
		doc, err := documentIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		poll := &datastruct.Poll{ID: doc.Ref.ID}
		doc.DataTo(poll)
		polls = append(polls, poll)
		i++
	}
	return polls, nil
}

func (e *electionQuery) UpdatePoll() *datastruct.Poll {
	return nil
}
func (e *electionQuery) DeletePoll(id common.ID_t) (common.ID_t, error) {
	return common.ID_t(""), common.ErrorNotImplemented
}
