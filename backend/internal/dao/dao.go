package repository

type DAO interface {
	NewVoteQuery() VoteQuery
	NewPollQuery() PollQuery
}
