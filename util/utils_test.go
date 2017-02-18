package util

import (
	"errors"
	"testing"
)

func TestNoErrorNoPanic(t *testing.T) {
	// does not throw
	CheckErr(nil)
}

func TestErrorPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	err := errors.New("test error")
	CheckErr(err)
}
