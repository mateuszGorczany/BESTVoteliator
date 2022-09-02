package rest

import (
	"net/http"

	gin "github.com/gin-gonic/gin"

	application "github.com/mateuszGorczany/BESTVoteliator/internal/app"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"
	"github.com/mateuszGorczany/BESTVoteliator/internal/service"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

func Run() {
	router := gin.Default()
	m, _ := application.NewMicroserviceApp()
	router.GET("/election", createElection(m.ElectingService))
	router.GET("/vote", getVote(m.VotingService))
	router.GET("/votes", getVotes(m.VotingService))

	router.Run()
}

func createElection(e service.ElectingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		vote, _ := e.CreateElection(dto.Election{})
		c.Bind(&dto.Election{})
		c.JSON(http.StatusOK, vote)
	}
}

func deleteElection(e service.ElectingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// vote, _ := e.
		election, _ := e.GetElection(common.ID_t(1234))
		e.DeleteElection(common.ID_t(1234))
		c.Bind(&dto.Election{})
		c.JSON(http.StatusOK, election)
	}
}

func getVote(v service.VotingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		vote, _ := v.GetVote(2134)
		c.Bind(&dto.Vote{})
		c.JSON(http.StatusOK, vote)
	}
}

func getVotes(v service.VotingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		votes, _ := v.GetVotes()
		c.Bind(&[]dto.Vote{{}, {}})
		c.JSON(http.StatusOK, votes)
	}
}
