package spreadsheets

import (
	"fmt"

	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
	"google.golang.org/api/sheets/v4"
)

type pollQuery struct {
	storage *Storage
}

func (e *pollQuery) CreatePoll(poll datastruct.Poll) (id common.ID_t, err error) {
	// e.storage.Service.Spreadsheets.Create(&sheets.Spreadsheet{})
	e.storage.Spreadsheets.Create(&sheets.Spreadsheet{
		Sheets: []*sheets.Sheet{&sheets.Sheet{}},
	}).Do()
	fmt.Printf("%v", poll)
	return common.ID_t(""), nil
}

func (e *pollQuery) GetPoll(id common.ID_t) (*datastruct.Poll, error) {
	return &datastruct.Poll{
		Name:        "Super",
		Description: "Super głosowanie",
		Options: []datastruct.Option{
			{
				Description: "option1",
				OptionID:    1,
			},
			{
				Description: "option1",
				OptionID:    2,
			},
		},
	}, nil
}

func (e *pollQuery) GetPolls() ([]*datastruct.Poll, error) {
	// values, _ := e.storage.Spreadsheets.Values.Get("id", "").Do()
	// values.Values
	return nil, nil
}

func (e *pollQuery) UpdatePoll() *datastruct.Poll {
	return nil
}

func (e *pollQuery) DeletePoll(id common.ID_t) (common.ID_t, error) {
	return "", common.ErrorNotImplemented
}
