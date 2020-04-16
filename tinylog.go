package tinylog

import (
	"fmt"
	"log"
	"os"

	f "github.com/lynxsecurity/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type Tiny struct {
	*logrus.Logger
}

func New(output *os.File) *Tiny {
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel)
	log.SetOutput(output)
	l.SetFormatter(&f.Formatter{
		NoColors:    false,
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
	return &Tiny{l}
}
func (t *Tiny) NewWarning(message string) {
	t.WithFields(logrus.Fields{"host": os.Getenv("HOSTNAME")}).Warn(message)
}
func (t *Tiny) ConsumerReady() {
	t.WithFields(logrus.Fields{"host": os.Getenv("HOSTNAME")}).Info(fmt.Sprintf("Consumer ready, PID: %d", os.Getpid()))
}
func (t *Tiny) Info(message string) {
	t.WithFields(logrus.Fields{"host": os.Getenv("HOSTNAME")}).Info(message)
}
