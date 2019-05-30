package workers

import (
	"fmt"
	"time"

	"gitlab.com/patricksangian/go-rest-mux/helpers/logger"
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
