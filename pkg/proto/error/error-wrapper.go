// Copyright (c) 2016 The Upspin Authors. All rights reserved.
// A lot of this was stolen from upspin
// (https://github.com/upspin/upspin/blob/master/errors/errors.go) and then
// modified.

package error

// populateStack uses the runtime to populate the Error's stack struct with
import (
	"bytes"
	"fmt"
	"runtime"
)

type IError interface {
	GetError() *Error
}

// assert Error implements the error interface.
var _ error = &Error{}

// Error implements the error interface.
func (e *Error) Error() string {
	b := new(bytes.Buffer)

	pad(b, ": ")
	b.WriteString(e.Message)
	if b.Len() == 0 {
		return "no error"
	}
	return b.String()
}

// WrapErr returns a corev1.Error for the given error and msg.
func WrapErr(err error, msg string) error {
	if err == nil {
		return nil
	}
	e := &Error{Message: fmt.Sprintf("%s: %s", msg, err.Error())}

	return e
}

// E is a useful func for instantiating corev1.Errors.
func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to E with no arguments")
	}
	e := &Error{}
	b := new(bytes.Buffer)
	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			pad(b, ": ")
			b.WriteString(arg)
		case int32:
			e.Code = arg
		case int:
			e.Code = int32(arg)
		case error:
			pad(b, ": ")
			b.WriteString(arg.Error())
		}
	}
	e.Message = b.String()

	return e
}

// frame returns the nth frame, with the frame at top of stack being 0.
func frame(callers []uintptr, n int) *runtime.Frame {
	frames := runtime.CallersFrames(callers)
	var f runtime.Frame
	for i := len(callers) - 1; i >= n; i-- {
		var ok bool
		f, ok = frames.Next()
		if !ok {
			break // Should never happen, and this is just debugging.
		}
	}
	return &f
}

// callers is a wrapper for runtime.callers that allocates a slice.
func callers() []uintptr {
	var stk [64]uintptr
	const skip = 4 // Skip 4 stack frames; ok for both E and Error funcs.
	n := runtime.Callers(skip, stk[:])
	return stk[:n]
}

var separator = ":\n\t"

// pad appends str to the buffer if the buffer already has some data.
func pad(b *bytes.Buffer, str string) {
	if b.Len() == 0 {
		return
	}
	b.WriteString(str)
}
