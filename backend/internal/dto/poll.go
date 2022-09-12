package dto

import common "github.com/mateuszGorczany/BESTVoteliator/utils"

type Poll struct {
	ID          common.ID_t `json:"id,omitempty"`
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Options     []Option    `json:"options" validate:"required"`
}

type Option struct {
	OptionID    int    `json:"option_id" validate:"required"`
	Description string `json:"description" validate:"required"`
}
