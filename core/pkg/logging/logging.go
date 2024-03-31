package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(logrus.InfoLevel)
}

// Info logs informational messages
func Info(message string) {
	log.Info(message)
}

// Warning logs warning messages
func Warning(message string) {
	log.Warning(message)
}

// Error logs error messages
func Error(message string) {
	log.Error(message)
}

// Fatal logs fatal messages and exits the application
func Fatal(message string) {
	log.Fatal(message)
}
