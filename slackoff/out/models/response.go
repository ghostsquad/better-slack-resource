package models

import (
	commonModels "github.com/ghostsquad/slack-off/common/models"
)

type Response struct {
	Version commonModels.Version        	`json:"version"`
	Metadata []commonModels.MetadataPair	`json:"metadata"`
}
