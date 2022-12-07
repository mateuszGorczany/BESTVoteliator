package dto

type Vote struct {
	ID        string `json:"id,omitempty"`
	PollID    string `json:"poll_id" validate:"required"`
	OptionID  string `json:"option_id" validate:"required"`
	UserID    string `json:"user,omitempty"`
	Timestamp string `json:"Timestamp,omitempty"`
}
