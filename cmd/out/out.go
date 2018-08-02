package main

import (
	"os"
	"encoding/json"
	"github.com/ghostsquad/slack-off"
	outModels "github.com/ghostsquad/slack-off/out/stepmodels"
)

func main() {
	if len(os.Args) < 2 {
		slackoff.Sayf("usage: %s <sources directory>\n", os.Args[0])
		os.Exit(1)
	}

	sourceDir := os.Args[1]

	if stat, err := os.Stat(sourceDir); err != nil || stat.IsDir() {
		slackoff.Fatal("provided source directory either doesn't exist, or is not a directory!", err)
	}

	//request := readRequest()

	//response, err := Run(sourceDir, request)
	//if err != nil {
	//	slackoff.Fatal("running command", err)
	//}

	//outputResponse(response)
}

func readRequest() *outModels.Request {
	request := &outModels.Request{}

	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		slackoff.Fatal("reading request from stdin", err)
	}

	return request
}

func outputResponse(response outModels.Response) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		slackoff.Fatal("writing response to stdout", err)
	}
}
