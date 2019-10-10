package formatter

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// StackdriverLogEntry ...
type StackdriverLogEntry struct {
	Severity string                 `json:"severity"`
	Time     json.Number            `json:"time"`
	Message  map[string]interface{} `json:"message,omitempty"`
}

// Stackdriver ...
type Stackdriver struct {
}

// Fotmat returns ...
func (*Stackdriver) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	fmt.Println(time.Now())
}
