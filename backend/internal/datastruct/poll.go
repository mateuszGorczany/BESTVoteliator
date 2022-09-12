package datastruct

type Poll struct {
	ID          string
	Name        string
	Description string   `json:"description"`
	Options     []Option `json:"options"`
}

type Option struct {
	OptionID    int    `json:"option_id"`
	Description string `json:"description"`
}
