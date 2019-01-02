package main

import (
	"os"
	"github.com/ghostsquad/slack-off"
	outModels "github.com/ghostsquad/slack-off/out/stepmodels"
	"github.com/ghostsquad/slack-off/out"
	"encoding/json"
)

func main() {
	srcDir := getSourceDir()

	ioFileReader := &slackoff.IOFileReader{}
	request := outModels.NewRequest(srcDir, ioFileReader)

	err := json.NewDecoder(os.Stdin).Decode(request)
	reportAndExitAsNecessary(err)

	httpClient := &slackoff.HttpClient{}
	templatizer, err := out.NewTemplatizer(srcDir, request.Params, ioFileReader)
	reportAndExitAsNecessary(err)

	command := out.NewCommand(templatizer, os.Stderr, httpClient)

	response, err := command.Run(request)
	reportAndExitAsNecessary(err)

	response.Write(os.Stdout)
}

func getSourceDir() string {
	if len(os.Args) < 2 {
		printErrorMessage("Error: usage: %s <sources directory>", os.Args[0])
		os.Exit(1)
	}

	srcDir := os.Args[1]
	fi, err := os.Stat(srcDir)
	if os.IsNotExist(err) {
		printErrorMessage("Error: sources directory (%s) does not exist", srcDir)
		os.Exit(1)
	}

	if !fi.IsDir() {
		printErrorMessage("Error: sources (%s) is not a directory", srcDir)
		os.Exit(1)
	}

	return os.Args[1]
}

func reportAndExitAsNecessary(err error) {
	if err != nil {
		printErrorMessage("Error: %s", err)
		os.Exit(1)
	}
}

func printErrorMessage(format string, a ...interface{}) {
	os.Stderr.Write([]byte(slackoff.ErrorColor.Sprintf(format, a...)))
}
