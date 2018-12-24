package must

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// code copied from assertions
func didPanic(f assert.PanicTestFunc) (bool, interface{}) {

	didPanic := false
	var message interface{}
	func() {
		defer func() {
			if message = recover(); message != nil {
				didPanic = true
			}
		}()

		f()
	}()

	return didPanic, message

}
func TestArgNotNilPositive(t *testing.T) {
	panics, err := didPanic(func() {
		var val interface{}
		ArgNotNil(val, "val")
	})
	if !panics {
		assert.Fail(t, "method not panics")
		return
	}

	msg := fmt.Sprint(err)
	assert.Equal(t, "argument val is nil", msg)
}
func TestArgNotNilNegative(t *testing.T) {
	assert.NotPanics(t, func() {
		ArgNotNil("hello", "val")
	})
}

func TestArgBeValidNegative(t *testing.T) {
	assert.NotPanics(t, func() {
		ArgBeValid(true, "arg1")
	})
}
func TestArgBeValidPositive(t *testing.T) {
	panics, err := didPanic(func() {
		val := 0
		ArgBeValid(val > 1, "val invalid: (%d > 1) == false", val)
	})
	if !panics {
		assert.Fail(t, "method not panics")
		return
	}
	msg := fmt.Sprint(err)
	assert.Equal(t, "val invalid: (0 > 1) == false", msg)
}

func TestArgBeValidWithoutArgs(t *testing.T) {
	panics, err := didPanic(func() {
		val := 0
		ArgBeValid(val > 1, "val invalid")
	})
	if !panics {
		assert.Fail(t, "method not panics")
		return
	}
	msg := fmt.Sprint(err)
	assert.Equal(t, "val invalid", msg)
}
