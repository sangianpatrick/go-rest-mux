package eventemitter

import (
	"fmt"

	"gitlab.com/patricksangian/go-rest-mux/src/workers"
	"gopkg.in/gomail.v2"
)

// Emitter contains channels and behavior for emitting listener
type Emitter struct {
	printCH chan<- interface{}
	emailCH chan<- *gomail.Message
}

// NewEventEmitter initiate event emitter
func NewEventEmitter() Emitter {
	printCH := make(chan interface{})
	emailCH := make(chan *gomail.Message)
	go workers.OnPrint(printCH)
	go workers.OnSendEmail(emailCH)
	return Emitter{
		printCH: printCH,
		emailCH: emailCH,
	}
}

// EmitPrint will print data
func (e Emitter) EmitPrint(data interface{}) {
	ch := e.printCH
	ch <- data
}

// EmitEmailSender will send email
func (e Emitter) EmitEmailSender(senderName, senderEmail, subject, messageBody string, recipientEmail []string) {
	ch := e.emailCH
	messageTemplate := `
	<!DOCTYPE html>
	<html>
	<body>
	` + messageBody + `
	</body>
	</html>
	`
	messageProperty := gomail.NewMessage()
	messageProperty.SetHeader("From", fmt.Sprintf("%s <%s>", senderName, senderEmail))
	messageProperty.SetHeader("To", recipientEmail...)
	messageProperty.SetHeader("Subject", subject)
	messageProperty.SetBody("text/html", messageTemplate)
	ch <- messageProperty
}
