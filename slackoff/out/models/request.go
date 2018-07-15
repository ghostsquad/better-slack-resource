package models

import commonModels "github.com/ghostsquad/slack-off/common/models"

type Request struct {
	Source commonModels.Source	`json:"source"`
	Params Params         			`json:"params"`
}
