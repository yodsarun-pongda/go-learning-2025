package config

import log "github.com/sirupsen/logrus"

func InitLogConfig() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
