package stepmodels

import (
	"github.com/ghostsquad/slack-off/resourcemodels"
)

type Request struct {
	Source resourcemodels.Source	`json:"source"`
	Params Params         			`json:"params"`
}
