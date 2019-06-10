package eventemitter

import (
	"fmt"

	article "gitlab.com/patricksangian/go-rest-mux/src/modules/article"
	articleModel "gitlab.com/patricksangian/go-rest-mux/src/modules/article/model"
	"gitlab.com/patricksangian/go-rest-mux/src/workers"
	"gopkg.in/gomail.v2"
)

// Emitter contains channels and behavior for emitting listener
type Emitter struct {
	printCH         chan<- interface{}
	emailCH         chan<- *gomail.Message
	createArticleCH chan<- *articleModel.Article
	handler         map[string]interface{}
}

// NewEventEmitter initiate event emitter
func NewEventEmitter(handler map[string]interface{}) Emitter {
	printCH := make(chan interface{})
	emailCH := make(chan *gomail.Message)
	createArticleCH := make(chan *articleModel.Article)

	go workers.OnCreateArticle(handler["article"].(article.EventHandler), createArticleCH)
	go workers.OnPrint(printCH)
	go workers.OnSendEmail(emailCH)
	return Emitter{
		printCH:         printCH,
		emailCH:         emailCH,
		createArticleCH: createArticleCH,
		handler:         handler,
	}
}

// EmitCreateArticle will create new article
func (e Emitter) EmitCreateArticle(newArticle *articleModel.Article) {
	ch := e.createArticleCH
	ch <- newArticle
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
