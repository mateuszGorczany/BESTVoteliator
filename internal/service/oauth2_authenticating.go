package service

import (
	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
)

type OAuth2AuthenticationService interface {
	Login() error
}

type oAuth2AuthenticationService struct {
	repository repository.DAO
}

func NewOAuth2AuthenticationService(repo repository.DAO) OAuth2AuthenticationService {
	return &oAuth2AuthenticationService{repository: repo}
}

func (o *oAuth2AuthenticationService) Login() error {
	return nil
}
