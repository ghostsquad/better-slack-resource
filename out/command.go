package out

import (
	"github.com/ghostsquad/slack-off/resourcemodels"
	"github.com/ghostsquad/slack-off/out/stepmodels"
)

//
func Run(string, request stepmodels.Request) (*stepmodels.Response, error) {


	return &stepmodels.Response{
		Version:  resourcemodels.Version{},
		Metadata: []resourcemodels.MetadataPair{},
	}, nil
}
