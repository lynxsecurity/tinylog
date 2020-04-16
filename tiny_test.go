package tinylog

import (
	"os"
	"testing"
)

func TestTinyLog(t *testing.T) {
	l := New(os.Stdout)
	l.NewWarning("Error creating file")
	l.ConsumerReady()
	l.Info("This is an informative message")
}
