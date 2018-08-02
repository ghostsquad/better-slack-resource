package stepmodels

import (
	"github.com/ghostsquad/slack-off/resourcemodels"
)

type Response struct {
	Version resourcemodels.Version        	`json:"version"`
	Metadata []resourcemodels.MetadataPair	`json:"metadata"`
}
