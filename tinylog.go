package tinylog

import (
	"log"
	"os"

	f "github.com/lynxsecurity/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type Tiny struct {
	*logrus.Logger
}

func New() *Tiny {
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
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
func (t *Tiny) Ready(message string) {
	t.WithFields(logrus.Fields{"host": os.Getenv("HOSTNAME"), "pid": os.Getpid()}).Info(message)
}
func (t *Tiny) Info(message string) {
	t.WithFields(logrus.Fields{"host": os.Getenv("HOSTNAME")}).Info(message)
}
