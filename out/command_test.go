// +build unit

package out

import (
	"testing"
	"github.com/ghostsquad/slack-off"
	"github.com/tylerb/is"
	"reflect"
	"github.com/golang/mock/gomock"
	"github.com/ghostsquad/slack-off/mocks"
	"io"
	"bytes"
	"github.com/ghostsquad/slack-off/out/stepmodels"
	"github.com/ghostsquad/slack-off/resourcemodels"
)

type testable struct {
	command        *command
	srcDir	       string
	mockCtrl       *gomock.Controller
	mockFileReader slackoff.FileReader
	mockPoster     slackoff.HttpPoster
	buffer         io.Writer
}

// setup mocks and objects we can use for testing
func commandTestable(t *testing.T) (ta *testable) {
	mockCtrl := gomock.NewController(t)

	ta = &testable{
		srcDir: "/src",
		mockCtrl: mockCtrl,
		mockFileReader: mock_slack_off.NewMockFileReader(mockCtrl),
		mockPoster: mock_slack_off.NewMockHttpPoster(mockCtrl),
		buffer: new(bytes.Buffer),
	}

	ta.command = NewCommand(ta.srcDir, ta.mockFileReader, ta.buffer, ta.mockPoster)

	return
}

func TestNewCommand(t *testing.T) {
	is := is.New(t)

	// this is the only test where we aren't actually using the `command` in testable
	ct := commandTestable(t)
	defer ct.mockCtrl.Finish()

	c := NewCommand(ct.srcDir, ct.mockFileReader, ct.buffer, ct.mockPoster)

	is.Equal(c.srcDir, ct.srcDir)
	is.Equal(c.fileReader, ct.mockFileReader)
	is.Equal(c.writer, ct.buffer)
	is.Equal(c.httpPoster, ct.mockPoster)

	// this might not be necessary, as later tests should test behavior of the resulting object
	is.True(reflect.TypeOf(c) == reflect.TypeOf(&command{}))
}

func TestCommand_Run(t *testing.T) {
	is := is.New(t)

	ct := commandTestable(t)
	defer ct.mockCtrl.Finish()

	r := stepmodels.Request{
		Source: resourcemodels.Source{
			Url: "test url",
		},
		Params: stepmodels.Params{
			Template:
		}
	}

	resp, err := ct.command.Run()
}
