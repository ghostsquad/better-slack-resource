package stepmodels

import (
	"github.com/ghostsquad/slack-off/resourcemodels"
	"io"
	"encoding/json"
)

type Response struct {
	Version resourcemodels.Version        	`json:"version"`
	Metadata []resourcemodels.MetadataPair	`json:"metadata"`
}

func (res *Response) Write(re io.Writer) error {
	return json.NewEncoder(re).Encode(res)
}
