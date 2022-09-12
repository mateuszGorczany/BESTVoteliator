package rest

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	application "github.com/mateuszGorczany/BESTVoteliator/internal/app"
	"github.com/mateuszGorczany/BESTVoteliator/internal/dto"
	"github.com/mateuszGorczany/BESTVoteliator/internal/services"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
	"github.com/rs/cors"
)

type key int

const (
	requestIDKey key = 0
	claimsKey    key = 1
)

type requestID string

func (r requestID) toString() string {
	return string(r)
}

func Run() {
	m, _ := application.NewMicroserviceApp()
	_ = m
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},           // All origins
		AllowedMethods:   []string{"GET", "POST"}, // Allowing only get, just an example
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return tracing(logging(next))
	})

	voting := VotingController{m.VotingService}
	polling := PollingController{m.PollingService}
	api := router.PathPrefix("/api/v1").Subrouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(getFronendDir())))
	api.HandleFunc("/login", loginHandler).Methods("POST")
	// get current user
	api.HandleFunc("/user", authenticate(handleUser)).Methods("GET")
	// create a poll
	api.HandleFunc("/poll", authenticate(polling.CreatePoll)).Methods("POST", "OPTIONS")
	// get a poll
	api.HandleFunc("/poll/{id}", polling.GetPoll).Methods("GET")
	// vote in a poll
	api.HandleFunc("/poll/{poll_id}/vote", authenticate(voting.CreateVote)).Methods("GET")
	// get a vote in a poll
	api.HandleFunc("/poll/{poll_id}/vote/{vote_id}", authenticate(voting.GetVote)).Methods("GET")

	address := viper.GetString("ServiceAddress")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := address + ":" + port
	common.Logger.Info("starting BESTVoteliator HTTP Server")
	common.Logger.Info("server available", zap.String("host", host))
	http.ListenAndServe(host, c.Handler(router))
}

func getFronendDir() string {
	return path.Join(viper.GetString("PWD"), "frontend")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	type parameters struct {
		GoogleJWT string `json:"GoogleJWT" validate:"required"`
	}
	params := parameters{}
	err := common.JSONDecodeAndValidate(r.Body, &params)
	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Check if JWT is valid
	claims, err := services.ValidateGoogleJWT(params.GoogleJWT)
	if err != nil {
		respondWithError(w, http.StatusForbidden, fmt.Errorf("Invalid google auth, error: %v", err))
		return
	}
	user := dto.User{
		Firstname: claims.FirstName,
		Lastname:  claims.LastName,
		Email:     claims.Email,
		ID:        claims.Id,
	}
	respondWithJSON(w, http.StatusOK, user)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	user := dto.User{
		Firstname: claims.FirstName,
		Lastname:  claims.LastName,
		Email:     claims.Email,
		ID:        claims.Id,
	}
	respondWithJSON(w, http.StatusOK, user)
}
