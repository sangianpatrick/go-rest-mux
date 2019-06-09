package workers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gitlab.com/patricksangian/go-rest-mux/helpers/logger"
	"gopkg.in/gomail.v2"
)

// OnPrint is a listener
func OnPrint(ch <-chan interface{}) {
	logger.Info("OnPrint", "status is listening")
	for {
		select {
		case data := <-ch:
			time.Sleep(time.Second)
			fmt.Printf("OnPrint: %v\n", data)
		}
	}
}

// OnSendEmail is a listener
func OnSendEmail(ch <-chan *gomail.Message) {
	logger.Info("OnSendEmail", "status is listening")
	go func() {
		host := os.Getenv("EMAIL_HOST")
		port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
		username := os.Getenv("EMAIL_USERNAME")
		password := os.Getenv("EMAIL_PASSWORD")

		d := gomail.NewPlainDialer(host, port, username, password)

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-ch:
				logger.Info("OnSendEmail", fmt.Sprintf("Trying to send email to %v", m.GetHeader("To")[0]))
				if !ok {
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						logger.Error("OnSendEmail", err)
						continue
					}
					open = true
				}
				if err := gomail.Send(s, m); err != nil {
					logger.Error("OnSendEmail", err)
					continue
				}
				logger.Info("OnSendEmail", fmt.Sprintf("An email message has been sent to %v", m.GetHeader("To")[0]))
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						logger.Error("OnSendEmail", err)
					}
					open = false
				}
			}
		}
	}()
}
