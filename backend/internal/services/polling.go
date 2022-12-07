package services

import (
	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"

	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type PollingService interface {
	CreatePoll(poll dto.Poll) (common.ID_t, error)
	GetPoll(id common.ID_t) (*dto.Poll, error)
	GetPolls() ([]*dto.Poll, error)
	DeletePoll(id common.ID_t) (common.ID_t, error)
}

type pollingService struct {
	repository repository.DAO
}

func NewPollingService(repo repository.DAO) PollingService {
	return &pollingService{repository: repo}
}

func (e *pollingService) CreatePoll(poll dto.Poll) (id common.ID_t, err error) {
	options := make([]datastruct.Option, len(poll.Options))
	for i, option := range poll.Options {
		options[i] = datastruct.Option(option)
	}
	voteModel := datastruct.Poll{
		Name:        poll.Name,
		Description: poll.Description,
		Options:     []datastruct.Option(options),
	}
	return e.repository.NewPollQuery().CreatePoll(voteModel)
}

func (e *pollingService) GetPoll(id common.ID_t) (*dto.Poll, error) {
	pollModel, err := e.repository.NewPollQuery().GetPoll(id)
	if err != nil {
		return nil, err
	}
	poll := convertPollEntityToDTO(pollModel)
	poll.ID = id
	return poll, nil
}

func (e *pollingService) GetPolls() ([]*dto.Poll, error) {
	pollsModels, err := e.repository.NewPollQuery().GetPolls()
	if err != nil {
		return nil, err
	}
	polls := make([]*dto.Poll, len(pollsModels))
	for i, pollModel := range pollsModels {
		polls[i] = convertPollEntityToDTO(pollModel)
	}

	return polls, nil
}

func (e *pollingService) DeletePoll(id common.ID_t) (common.ID_t, error) {
	return "", common.ErrorNotImplemented
}

func convertPollEntityToDTO(pollModel *datastruct.Poll) *dto.Poll {
	options := make([]dto.Option, len(pollModel.Options))
	for i, option := range pollModel.Options {
		options[i] = dto.Option(option)
	}

	return &dto.Poll{
		ID:          common.ID_t(pollModel.ID),
		Name:        pollModel.Name,
		Description: pollModel.Description,
		Options:     options,
	}
}
