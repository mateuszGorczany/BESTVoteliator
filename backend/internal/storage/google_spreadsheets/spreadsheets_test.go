package spreadsheets

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func init() {
}

func TestClientConstruction(t *testing.T) {
	storage, _ := NewStorage()
	response, _ := storage.Spreadsheets.Values.Get(storage.SheetID, "A1:A2").Do()
	for _, row := range response.Values {
		for _, cell := range row {
			assert.Equal(t, cell, "")
		}
	}
}

func TestSavingData(t *testing.T) {
	storage, _ := NewStorage()
	// value = "my Value"
	// storage.Spreadsheets.Values.Append(storage.SheetID, "A3", &sheets.ValueRange{
	// 	Values: [][]interface{{"ok"}},
	// })
	response, _ := storage.Spreadsheets.Values.Get(storage.SheetID, "A1:A2").Do()
	for _, row := range response.Values {
		for _, cell := range row {
			assert.Equal(t, cell, "")
		}
	}
}

func Test_spreadsheets_test(t *testing.T) {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	readRange := "Class Data!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf("%s, %s\n", row[0], row[4])
		}
	}
}
