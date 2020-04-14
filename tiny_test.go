package tinylog

import (
	"testing"
)

func TestTinyLog(t *testing.T) {
	l := New()
	l.NewWarning("Error creating file")
	l.Ready("Consumer ready")
	l.Info("This is an informative message")
}
