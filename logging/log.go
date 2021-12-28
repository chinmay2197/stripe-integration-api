package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}

var Logger = log.New()
