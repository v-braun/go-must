// Package must is a simple assertion package for go
package must

import (
	"fmt"

	"github.com/pkg/errors"
)

// ArgNotNil panics if item is nil.
// The error will be 'argument {name} is nil'
func ArgNotNil(item interface{}, name string) {
	if item == nil {
		panic(errors.New(fmt.Sprintf("argument %s is nil", name)))
	}
}

// ArgBeValid checks the given expression
// and if it returns true it will panics with the given message and format args
func ArgBeValid(valid bool, message string, a ...interface{}) {
	if !valid {
		panic(errors.New(fmt.Sprintf(message, a...)))
	}
}
