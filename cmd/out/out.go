package main

import (
	"os"
	"github.com/ghostsquad/slack-off"
	outModels "github.com/ghostsquad/slack-off/out/stepmodels"
	"github.com/ghostsquad/slack-off/out"
)

func main() {
	request := outModels.Request{}
	err := request.Populate(os.Stdin)
	reportAndExitAsNecessary(err)

	httpClient := &slackoff.HttpClient{}
	ioFileReader := &slackoff.IOFileReader{}

	command := out.NewCommand(ioFileReader, os.Stderr, httpClient)

	response, err := command.Run(request)
	reportAndExitAsNecessary(err)

	response.Write(os.Stdout)
}

func reportAndExitAsNecessary(err error) {
	if err != nil {
		os.Stderr.Write([]byte(slackoff.ErrorColor.Sprintf("Error: %s", err)))
		os.Exit(1)
	}
}
