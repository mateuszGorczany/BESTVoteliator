package repository

type DAO interface {
	NewVoteQuery() VoteQuery
	NewElectionQuery() ElectionQuery
}
