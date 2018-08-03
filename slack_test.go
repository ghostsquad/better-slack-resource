package slackoff

import "testing"

func TestAssertSlackUrl(t *testing.T) {
	val := "https://hooks.slack.com/foo"

	err := AssertSlackUrl(val)

	if err != nil {
		t.Errorf("An error was received")
	}
}

//
//func TestAssertSlackUrlWrongSchema(t *testing.T) {
//	val := "http://hooks.slack.com/foo"
//
//	err := AssertSlackUrl(val)
//
//	if err == nil {
//		t.Errorf("An error was not received")
//	}
//}
//
//func TestAssertSlackUrlWrongDomain(t *testing.T) {
//	val := "https://notslack.com/foo"
//
//	err := AssertSlackUrl(val)
//
//	if err == nil {
//		t.Errorf("An error was not received")
//	}
//}
