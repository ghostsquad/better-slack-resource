package stepmodels

import (
"testing"
	"github.com/ghostsquad/slack-off"
	"github.com/tylerb/is"
	"github.com/golang/mock/gomock"
	"github.com/ghostsquad/slack-off/mocks"
)

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

	is.Msg("Validations did prevent error when both template and template_file are given, %s", err).NotNil(err)
}

func TestParams_RegisterValidations_WhenNeitherTemplateAndTemplateFileGiven(t *testing.T)  {
	is := is.New(t)

	p := Params{
	}

	val := slackoff.InitValidator()

	p.RegisterValidations(val)

	err := val.Struct(p)

	is.Msg("Validations did prevent error when neither template and template_file are given, %s").NotNil(err)
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
	mockReader.EXPECT().ReadFile(expectedPath).Return(expectedContents, nil)

	tmpl, err := p.GetTemplate(mockReader)

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

	tmpl, err := p.GetTemplate(mockReader)

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

	channels, err := p.GetExtraChannels(mockReader)

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

	channels, err := p.GetExtraChannels(mockReader)

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
	mockReader.EXPECT().ReadFile(expectedPath).Return("ch1    ch2\n ch3\tch4", nil)

	channels, err := p.GetExtraChannels(mockReader)

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
	mockReader.EXPECT().ReadFile(expectedPath).Return("ch3 ch4", nil)

	channels, err := p.GetExtraChannels(mockReader)

	is.Nil(err)
	is.Equal(channels, []string{"ch1", "ch2", "ch3", "ch4"})
}
