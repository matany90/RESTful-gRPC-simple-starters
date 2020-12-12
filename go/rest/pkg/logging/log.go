package logging

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// Setup function will init Logger
// global object
func Setup() {
	// create logger new instance
	Log = logrus.New()

	// set logger lovel to debug
	Log.Level = logrus.DebugLevel

	// set logger output format
	Log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}

	// set the logger output dest
	Log.Out = os.Stdout
}