package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"
	"github.com/mateuszGorczany/BESTVoteliator/internal/services"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type VotingController struct {
	votingService services.VotingService
}

func (vc *VotingController) CreateVote(w http.ResponseWriter, r *http.Request) {

	vote := dto.Vote{
		PollID:   mux.Vars(r)["poll_id"],
		UserID:   getClaims(r).Email,
		OptionID: r.URL.Query().Get("option"),
	}
	voteID, err := vc.votingService.CreateVote(vote)
	if err != nil {
		respondWithError(w, http.StatusConflict, err)
	}
	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"vote_id": voteID,
		"user_id": vote.UserID,
	})
}

func (vc *VotingController) GetVote(w http.ResponseWriter, r *http.Request) {
	pollID := mux.Vars(r)["poll_id"]
	_ = pollID
	voteID := mux.Vars(r)["vote_id"]
	if voteID == "all" {
		vc.GetVotes(w, r)
		return
	}
	vote, err := vc.votingService.GetVote(common.ID_t(voteID))
	if err != nil {
		respondWithError(w, http.StatusConflict, err)
		return
	}
	// if pollID != vote.PollID {
	// 	respondWithError(w, http.StatusBadRequest, fmt.Errorf("incorrect PollID"))
	// 	return
	// }
	respondWithJSON(w, http.StatusOK, vote)
}

func (vc *VotingController) GetVotes(w http.ResponseWriter, r *http.Request) {
	pollID := mux.Vars(r)["poll_id"]
	votes, err := vc.votingService.GetVotes(common.ID_t(pollID))
	if err != nil {
		respondWithError(w, http.StatusConflict, err)
	}
	if len(votes) > 0 && pollID != votes[0].PollID {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("incorrect PollID"))
	}
	respondWithJSON(w, http.StatusCreated, votes)
}
