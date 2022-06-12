package repository

type ElectionQuery interface {
	CreateElection()
	GetElection()
	GetElections()
	UpdateElection()
	DeleteElection()
}
