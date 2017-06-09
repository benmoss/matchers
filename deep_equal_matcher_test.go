package matchers_test

import (
	"testing"

	. "github.com/benmoss/matchers"
)

type Test struct {
	a, b    interface{}
	match   bool
	message string
}

func strp(x string) *string {
	return &x
}

var tests []Test

func init() {
	tests = []Test{
		{1, 1, true, ""},
		{1, 2, false, "Expected\n    <int>: 2\nto deeply equal\n    <int>: 1\nmismatch at top level: unequal; obtained 2; expected 1"},
	}
}

func TestMatch(t *testing.T) {
	for _, test := range tests {
		matcher := DeepEqual(test.a)
		match, _ := matcher.Match(test.b)
		if match != test.match {
			t.Errorf("fail comparing %v and %v. expected %v, got %v", test.a, test.b, test.match, match)
		}
	}
}

func TestFailureMessage(t *testing.T) {
	for _, test := range tests {
		matcher := DeepEqual(test.a)
		_, _ = matcher.Match(test.b)

		expectedMsg := test.message
		if expectedMsg == "" {
			continue
		}
		actualMsg := matcher.FailureMessage(test.b)
		if expectedMsg != actualMsg {
			t.Errorf("expected matcher message %q, got %q.", expectedMsg, actualMsg)
		}
	}
}
