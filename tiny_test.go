package tinylog

import (
	"errors"
	"testing"
)

func TestTinyLog(t *testing.T) {
	l := New()
	l.NewError(errors.New("critical error"), "Error creating file")
	l.Ready("Consumer ready")
	l.Info("This is an informative message")
}
