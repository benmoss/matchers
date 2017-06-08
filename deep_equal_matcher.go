package matchers

import (
	"fmt"

	"github.com/juju/testing/checkers"
)

// DeepEqual uses github.com/juju/testing's reimplementation of
// reflect.DeepEqual. Their docstring:
//
// DeepEqual tests for deep equality. It uses normal == equality where
// possible but will scan elements of arrays, slices, maps, and fields
// of structs. In maps, keys are compared with == but elements use deep
// equality. DeepEqual correctly handles recursive types. Functions are
// equal only if they are both nil.
//
// DeepEqual differs from reflect.DeepEqual in two ways:
// - an empty slice is considered equal to a nil slice.
// - two time.Time values that represent the same instant
// but with different time zones are considered equal.
//
// If the two values compare unequal, the resulting error holds the
// first difference encountered.
func DeepEqual(expected interface{}) *deepMatcher {
	return &deepMatcher{expected: expected}
}

type deepMatcher struct {
	expected    interface{}
	explanation error
}

func (d *deepMatcher) Match(actual interface{}) (bool, error) {
	match, err := checkers.DeepEqual(d.expected, actual)
	d.explanation = err
	return match, nil
}

func (d *deepMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected %v to deeply equal %v.\nReason: %q",
		d.expected, actual, d.explanation)
}

func (d *deepMatcher) NegatedFailureMessage(actual interface{}) string {
	return d.FailureMessage(actual)
}
