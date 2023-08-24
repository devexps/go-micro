package sse

import (
	"encoding/base64"
	"time"
)

// Event holds all the event source fields
type Event struct {
	timestamp time.Time
	ID        []byte
	Data      []byte
	Event     []byte
	Retry     []byte
	Comment   []byte
}

func (e *Event) hasContent() bool {
	return len(e.ID) > 0 || len(e.Data) > 0 || len(e.Event) > 0 || len(e.Retry) > 0
}

func (e *Event) encodeBase64() {
	dataLen := len(e.Data)
	if dataLen > 0 {
		output := make([]byte, base64.StdEncoding.EncodedLen(dataLen))
		base64.StdEncoding.Encode(output, e.Data)
		e.Data = output
	}
}
