package stepmodels

import (
	"testing"
	"github.com/tylerb/is"
	"bytes"
	"github.com/ghostsquad/slack-off/resourcemodels"
)

func TestResponse_Write(t *testing.T) {
	is := is.New(t)

	buf := new(bytes.Buffer)

	r := &Response{
		Version: resourcemodels.Version{
			Timestamp: 1,
		},
	}

	err := r.Write(buf)

	is.Nil(err)

	result := buf.String()
	is.Equal(result, `{"version":{"timestamp":1}}` + "\n")
}

