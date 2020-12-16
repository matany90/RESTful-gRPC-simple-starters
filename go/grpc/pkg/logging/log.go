package logging

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// define logger global instance
var Log *logrus.Logger

// Setup initialize the log instance
func Setup() {
	// create a new instance of the logrus logger
	Log = logrus.New()

	// set the logger level to debug
	Log.Level = logrus.DebugLevel

	// set the logger output format
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
