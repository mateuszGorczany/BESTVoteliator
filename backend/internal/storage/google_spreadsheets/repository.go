package spreadsheets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Storage struct {
	*sheets.Service
	SheetID string
}

func NewStorage() (*Storage, error) {
	// Service account based oauth2 two legged integration
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(""), option.WithScopes(sheets.SpreadsheetsScope))

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}
	c := &Storage{
		Service: srv,
		SheetID: viper.GetString("Database.Sheets.SheetID"),
	}

	return c, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func (s *Storage) NewVoteQuery() repository.VoteQuery {
	fmt.Print("New vote query ccalled from spreadsheet\n")
	return &voteQuery{storage: s}
}

func (s *Storage) NewPollQuery() repository.PollQuery {
	return &pollQuery{storage: s}
}
