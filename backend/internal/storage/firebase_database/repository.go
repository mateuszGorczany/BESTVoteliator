package database

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	"github.com/spf13/viper"

	"google.golang.org/api/option"
)

var (
	DB *firestore.Client
)

func PollQueryBuiler() *firestore.CollectionRef {
	return DB.Collection("poll")
}

func VoteQueryBuilder() *firestore.CollectionRef {
	return DB.Collection("vote")
}

type storage struct{}

func NewStorage() (repository.DAO, error) {
	credentialsPath := viper.GetString("Database.Firebase.Credentials")
	opt := option.WithCredentialsFile(credentialsPath)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %w", err)
	}
	DB, err = app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return &storage{}, nil
}

func (s *storage) NewVoteQuery() repository.VoteQuery {
	return &voteQuery{}
}

func (s *storage) NewPollQuery() repository.PollQuery {
	return &electionQuery{}
}
