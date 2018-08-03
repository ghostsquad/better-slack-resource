// +build integration

package main

import (
	"testing"
	"github.com/tylerb/is"
)

func Test(t *testing.T) {
	is := is.New(t)

	is.False(true)
}
