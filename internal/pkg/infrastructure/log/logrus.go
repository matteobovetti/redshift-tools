package log

import (
	"github.com/sirupsen/logrus"
)

var (
	QuietLevel           bool
	VerboseLevel         bool
	VeryVerboseLevel     bool
	VeryVeryVerboseLevel bool
)

func SetUp() error {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	setUpLogLevel(logrus.TraceLevel)

	return nil
}

func setUpLogLevel(defaultLoglevel logrus.Level) {
	if VeryVeryVerboseLevel {
		logrus.SetLevel(logrus.TraceLevel)

		return
	}

	if VeryVerboseLevel {
		logrus.SetLevel(logrus.DebugLevel)

		return
	}

	if VerboseLevel {
		logrus.SetLevel(logrus.InfoLevel)

		return
	}

	if QuietLevel {
		logrus.SetLevel(logrus.PanicLevel)

		return
	}

	logrus.SetLevel(defaultLoglevel)
}
