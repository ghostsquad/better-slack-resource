package stepmodels

import (
	"testing"
	"github.com/ghostsquad/slack-off"
	"github.com/tylerb/is"
	"github.com/golang/mock/gomock"
	"github.com/ghostsquad/slack-off/mocks"
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
)

func TestParams_Validations_WhenFileVarsEmpty(t *testing.T) {
	is := is.New(t)

	p := Params{}
	val := slackoff.InitValidator()
	err := val.Struct(p)

	is.Msg("empty vars should be ok, %s", err).Nil(err)
}

func TestParams_Validations_WhenFileVarsNotEmpty(t *testing.T) {
	is := is.New(t)

	p := Params{
		FileVars: map[string]string{
			"foo": "bar",
		},
	}
	val := slackoff.InitValidator()
	err := val.Struct(p)

	is.Msg("valid map of vars should be ok, %s", err).Nil(err)
}

func TestParams_Validations_WhenFileVarsEmptyKey(t *testing.T) {
	is := is.New(t)

	p := Params{
		FileVars: map[string]string{
			"": "bar",
		},
	}
	val := slackoff.InitValidator()
	err := val.Struct(p)

	is.Msg("Validations did not prevent an empty key, %s", err).NotNil(err)
}

func TestParams_Validations_WhenFileVarsEmptyValue(t *testing.T) {
	is := is.New(t)

	p := Params{
		FileVars: map[string]string{
			"foo": "",
		},
	}
	val := slackoff.InitValidator()
	err := val.Struct(p)

	is.Msg("Validations did not prevent an empty value, %s", err).NotNil(err)
}

func TestParams_Validations_VarsEmpty(t *testing.T) {
	is := is.New(t)

	p := Params{}
	val := slackoff.InitValidator()
	err := val.Struct(p)

	is.Msg("empty vars should be ok, %s", err).Nil(err)
}

func TestParams_Validations_VarsFull(t *testing.T) {
	is := is.New(t)

	p := Params{
		Vars: map[string]string{
			"foo": "bar",
		},
	}
	val := slackoff.InitValidator()
	err := val.Struct(p)

	is.Msg("full vars should be ok, %s", err).Nil(err)
}

func TestParams_Validations_VarsEmptyKey(t *testing.T) {
	is := is.New(t)

	p := Params{
		Vars: map[string]string{
			"": "bar",
		},
	}
	val := slackoff.InitValidator()
	err := val.Struct(p)

	is.Msg("validation did not prevent an empty key, %s", err).NotNil(err)
}

func TestParams_Validations_VarsEmptyValue(t *testing.T) {
	is := is.New(t)

	p := Params{
		Vars: map[string]string{
			"bar": "",
		},
	}
	val := slackoff.InitValidator()
	err := val.Struct(p)

	is.Msg("validation did not prevent an empty value, %s", err).NotNil(err)
}

func TestParams_RegisterValidations_WhenTemplateGiven(t *testing.T) {
	is := is.New(t)

	p := Params{
		Template: "foo",
	}

	val := slackoff.InitValidator()

	p.RegisterValidations(val)

	err := val.Struct(p)

	is.Msg("Validations failed, %s", err).Nil(err)
}

func TestParams_RegisterValidations_WhenTemplateFileGiven(t *testing.T)  {
	is := is.New(t)

	p := Params{
		TemplateFile: "foo",
	}

	val := slackoff.InitValidator()

	p.RegisterValidations(val)

	err := val.Struct(p)

	is.Msg("Validations failed, %s", err).Nil(err)
}

func TestParams_RegisterValidations_WhenTemplateAndTemplateFileGiven(t *testing.T)  {
	is := is.New(t)

	p := Params{
		Template: "foo",
		TemplateFile: "foo",
	}

	val := slackoff.InitValidator()

	p.RegisterValidations(val)

	err := val.Struct(p)

	is.Msg("Validations did not prevent error when both template and template_file are given, %s", err).NotNil(err)
}

func TestParams_RegisterValidations_WhenNeitherTemplateAndTemplateFileGiven(t *testing.T)  {
	is := is.New(t)

	p := Params{
	}

	val := slackoff.InitValidator()

	p.RegisterValidations(val)

	err := val.Struct(p)

	is.Msg("Validations did not prevent error when neither template and template_file are given, %s").NotNil(err)
}

func TestParams_GetTemplate_WhenTemplateFileGiven(t *testing.T) {
	is := is.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedPath := "path/to/template"
	expectedContents := "test contents"

	p := Params{
		TemplateFile: expectedPath,
	}

	mockReader := mock_slack_off.NewMockFileReader(mockCtrl)
	mockReader.EXPECT().ReadFile("/src/" + expectedPath).Return(expectedContents, nil)

	tmpl, err := p.GetTemplate("/src", mockReader)

	is.Nil(err)
	is.Equal(tmpl, expectedContents)
}

func TestParams_GetTemplate_WhenTemplateGiven(t *testing.T) {
	is := is.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedContents := "test contents"

	p := Params{
		Template: expectedContents,
	}

	mockReader := mock_slack_off.NewMockFileReader(mockCtrl)

	tmpl, err := p.GetTemplate("/src", mockReader)

	is.Nil(err)
	is.Equal(tmpl, expectedContents)
}

func TestParams_GetExtraChannels_WhenChannelAppendHasOneValue(t *testing.T) {
	is := is.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	p := Params{
		ChannelAppend: "ch1",
	}

	mockReader := mock_slack_off.NewMockFileReader(mockCtrl)

	channels, err := p.GetExtraChannels("/src", mockReader)

	is.Nil(err)
	is.Equal(channels, []string{"ch1"})
}

func TestParams_GetExtraChannels_WhenChannelAppendHasMultipleValue(t *testing.T) {
	is := is.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	p := Params{
		ChannelAppend: "ch1    ch2",
	}

	mockReader := mock_slack_off.NewMockFileReader(mockCtrl)

	channels, err := p.GetExtraChannels("/src", mockReader)

	is.Nil(err)
	is.Equal(channels, []string{"ch1", "ch2"})
}

func TestParams_GetExtraChannels_WhenChannelFileIncluded(t *testing.T) {
	is := is.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedPath := "path/to/channels"

	p := Params{
		ChannelFile: expectedPath,
	}

	mockReader := mock_slack_off.NewMockFileReader(mockCtrl)
	mockReader.EXPECT().ReadFile("/src/" + expectedPath).Return("ch1    ch2\n ch3\tch4", nil)

	channels, err := p.GetExtraChannels("/src", mockReader)

	is.Nil(err)
	is.Equal(channels, []string{"ch1", "ch2", "ch3", "ch4"})
}

func TestParams_GetExtraChannels_WhenChannelAppendAndChannelFileIncluded(t *testing.T) {
	is := is.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedPath := "path/to/channels"

	p := Params{
		ChannelAppend: "ch1 ch2",
		ChannelFile: expectedPath,
	}

	mockReader := mock_slack_off.NewMockFileReader(mockCtrl)
	mockReader.EXPECT().ReadFile("/src/" + expectedPath).Return("ch3 ch4", nil)

	channels, err := p.GetExtraChannels("/src/", mockReader)

	is.Nil(err)
	is.Equal(channels, []string{"ch1", "ch2", "ch3", "ch4"})
}

func TestParams_GetFileVars(t *testing.T) {
	is := is.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedPath := "path/to/file1"
	expectedKey := "fookey"
	expectedValue := "foovalue"

	p := Params{
		FileVars: map[string]string{
			expectedKey: expectedPath,
		},
	}

	mockReader := mock_slack_off.NewMockFileReader(mockCtrl)
	mockReader.EXPECT().ReadFile("/src/" + expectedPath).Return(expectedValue, nil)

	fileVars, err := p.GetFileVars("/src", mockReader)

	is.Nil(err)
	is.Equal(fileVars, map[string]string{expectedKey: expectedValue})
}

func TestParams_GetFileVars_WhenMultipleFiles(t *testing.T) {
	is := is.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedPath1 := "path/to/file1"
	expectedKey1 := "fookey1"
	expectedValue1 := "foovalue1"

	expectedPath2 := "path/to/file2"
	expectedKey2 := "fookey2"
	expectedValue2 := "foovalue2"

	p := Params{
		FileVars: map[string]string{
			expectedKey1: expectedPath1,
			expectedKey2: expectedPath2,
		},
	}

	mockReader := mock_slack_off.NewMockFileReader(mockCtrl)
	mockReader.EXPECT().ReadFile("/src/" + expectedPath1).Return(expectedValue1, nil)
	mockReader.EXPECT().ReadFile("/src/" + expectedPath2).Return(expectedValue2, nil)

	fileVars, err := p.GetFileVars("/src", mockReader)

	is.Nil(err)
	is.Equal(fileVars, map[string]string{
		expectedKey1: expectedValue1,
		expectedKey2: expectedValue2,
	})
}

func TestParams_GetFileVars_WhenReadErrorOccurs(t *testing.T) {
	is := is.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedPath1 := "path/to/file1"
	expectedKey1 := "fookey1"

	expectedPath2 := "path/to/file2"
	expectedKey2 := "fookey2"

	p := Params{
		FileVars: map[string]string{
			expectedKey1: expectedPath1,
			expectedKey2: expectedPath2,
		},
	}

	mockReader := mock_slack_off.NewMockFileReader(mockCtrl)
	mockReader.EXPECT().ReadFile("/src/" + expectedPath1).Return("", errors.New(fmt.Sprintf("file read err: %s", expectedPath1)))
	mockReader.EXPECT().ReadFile("/src/" + expectedPath2).Return("", errors.New(fmt.Sprintf("file read err: %s", expectedPath2)))

	_, err := p.GetFileVars("/src", mockReader)

	is.NotNil(err)
	mErr, ok := err.(*multierror.Error)
	is.True(ok)
	is.Len(mErr.Errors, 2)
}
