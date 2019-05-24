package logger

import (
	"fmt"
	"log"
	"os"

	raven "github.com/getsentry/raven-go"
)

func ravenInit() {
	raven.SetDSN(os.Getenv("SENTRY_URL"))
}

// Fatal logger
func Fatal(ctx string, err error) {
	ravenInit()
	captured := fmt.Sprintf(`[FATAL] %s: %s`, ctx, err.Error())
	raven.CaptureMessageAndWait(captured, nil)
	log.Fatalln(captured)
}

// Error logger
func Error(ctx string, err error) {
	go func() {
		ravenInit()
		captured := fmt.Sprintf(`[ERROR] %s: %s`, ctx, err.Error())
		raven.CaptureMessage(captured, nil)
		log.Println(captured)
	}()
}

// Info logger
func Info(ctx string, message string) {
	go func() {
		ravenInit()
		captured := fmt.Sprintf(`[INFO] %s: %s`, ctx, message)
		raven.CaptureMessage(captured, nil)
		log.Println(captured)
	}()
}
