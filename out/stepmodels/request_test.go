// +build unit

package stepmodels

import (
	"testing"
	"github.com/tylerb/is"
	"strings"
	"github.com/ghostsquad/slack-off"
	"github.com/ghostsquad/slack-off/resourcemodels"
)

func TestRequest_Populate(t *testing.T) {
	is := is.New(t)

	contents := `{"params":{"Template": "Foo"}}`

	r := &Request{}
	err := r.Populate(strings.NewReader(contents))

	is.Nil(err)
	is.Equal(r.Params.Template, "Foo")
}

func TestRequest_Populate_WhenInvalidJson(t *testing.T) {
	is := is.New(t)

	contents := `{"params":`

	r := &Request{}
	err := r.Populate(strings.NewReader(contents))

	is.NotNil(err)
}

func TestRequest_Validations_WhenEmpty(t *testing.T)  {
	is := is.New(t)

	val := slackoff.InitValidator()
	r := Request{}
	err := val.Struct(r)

	is.Msg("Validations should prevent empty fields, %s", err).NotNil(err)
}

func TestRequest_Validations_WhenSourceEmpty(t *testing.T)  {
	is := is.New(t)

	val := slackoff.InitValidator()
	r := Request{
		Params: Params{
			Template: "foo",
		},
	}
	err := val.Struct(r)

	is.Msg("Validations should prevent empty source field, %s", err).NotNil(err)
}

func TestRequest_Validations_WhenValidationsRegistered_WhenParamsEmpty(t *testing.T)  {
	is := is.New(t)

	val := slackoff.InitValidator()
	r := Request{
		Source: resourcemodels.Source{
			Url: "foo",
		},
	}
	r.Params.RegisterValidations(val)
	err := val.Struct(r)

	is.Msg("Validations should prevent empty params field, %s", err).NotNil(err)
}

func TestRequest_Validations_WhenFull(t *testing.T)  {
	is := is.New(t)

	val := slackoff.InitValidator()
	r := Request{
		Params: Params{
			Template: "foo",
		},
		Source: resourcemodels.Source{
			Url: "foo",
		},
	}
	err := val.Struct(r)

	is.Msg("Validations should pass if everything is filled in, %s", err).Nil(err)
}
