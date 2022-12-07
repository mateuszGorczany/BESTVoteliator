package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"
	"github.com/mateuszGorczany/BESTVoteliator/internal/services"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
)

type PollingController struct {
	pollingService services.PollingService
}

func (p *PollingController) CreatePoll(w http.ResponseWriter, r *http.Request) {
	poll := dto.Poll{}
	err := common.JSONDecodeAndValidate(r.Body, &poll)
	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	pollID, err := p.pollingService.CreatePoll(poll)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	}
	respondWithJSON(w, http.StatusCreated, struct {
		PollID common.ID_t `json:"poll_id"`
	}{PollID: pollID})
}

func (p *PollingController) GetPoll(w http.ResponseWriter, r *http.Request) {
	pollID := mux.Vars(r)["id"]
	if pollID == "all" {
		p.GetPolls(w, r)
		return
	}
	poll, err := p.pollingService.GetPoll(common.ID_t(pollID))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	}
	respondWithJSON(w, http.StatusOK, poll)
}

func (p *PollingController) GetPolls(w http.ResponseWriter, r *http.Request) {
	polls, err := p.pollingService.GetPolls()
	fmt.Println(err)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	}
	respondWithJSON(w, http.StatusOK, polls)
}
