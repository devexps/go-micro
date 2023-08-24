package sse

import "net/url"

// Subscriber holds subscriber properties
type Subscriber struct {
	quit       chan *Subscriber
	connection chan *Event
	removed    chan struct{}
	eventId    int
	URL        *url.URL
}

// close will let the stream know that the clients connection has terminated
func (s *Subscriber) close() {
	s.quit <- s
	if s.removed != nil {
		<-s.removed
	}
}
