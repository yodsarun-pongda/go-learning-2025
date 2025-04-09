package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func InitLogConfig() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func Info(msg any) {
	log.Info(fmt.Sprintf("%v", msg))
}
func Debug(msg any) {
	log.Debug(fmt.Sprintf("%v", msg))
}
func Error(msg any, error any) {
	log.Error(fmt.Sprintf("%v, %v", msg, error))
}
func Fatal(msg any) {
	log.Fatal(fmt.Sprintf("%v", msg))
}
func Panic(msg any) {
	log.Panic(fmt.Sprintf("%v", msg))
}
func Trace(msg any) {
	log.Trace(fmt.Sprintf("%v", msg))
}
func Warn(msg any) {
	log.Warn(fmt.Sprintf("%v", msg))
}
