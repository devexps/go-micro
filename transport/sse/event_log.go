package sse

import (
	"strconv"
	"time"
)

type EventLog []*Event

// Add .
func (e *EventLog) Add(ev *Event) {
	if !ev.hasContent() {
		return
	}
	ev.ID = []byte(e.currentIndex())
	ev.timestamp = time.Now()
	*e = append(*e, ev)
}

// Clear .
func (e *EventLog) Clear() {
	*e = nil
}

// Replay .
func (e *EventLog) Replay(s *Subscriber) {
	for i := 0; i < len(*e); i++ {
		id, _ := strconv.Atoi(string((*e)[i].ID))
		if id >= s.eventId {
			s.connection <- (*e)[i]
		}
	}
}

func (e *EventLog) currentIndex() string {
	return strconv.Itoa(len(*e))
}
