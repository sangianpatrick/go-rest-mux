package eventsource

import (
	"github.com/sangianpatrick/go-rest-mux/src/app"
	"gopkg.in/mgo.v2"
)

// NewEventSource func
func NewEventSource(mgoSESS *mgo.Session) map[string]interface{} {
	eventSource := map[string]interface{}{
		"article": app.MountArticleEventSource(mgoSESS),
	}
	return eventSource
}
