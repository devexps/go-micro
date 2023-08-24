package sse

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// ServeHTTP serves new connections with events for a given stream ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	flusher, err := w.(http.Flusher)
	if !err {
		writeError(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	s.prepareHeaderForSSE(w)

	// Get the StreamID from the URL
	streamID := r.URL.Query().Get("stream")
	if streamID == "" {
		writeError(w, "Please specify a stream!", http.StatusInternalServerError)
		return
	}
	stream := s.streamMgr.Get(StreamID(streamID))
	if stream == nil {
		if !s.autoStream {
			writeError(w, "Stream not found!", http.StatusInternalServerError)
			return
		}
		stream = s.CreateStream(StreamID(streamID))
	}
	eventId := 0
	if id := r.Header.Get("Last-Event-ID"); id != "" {
		var err error
		eventId, err = strconv.Atoi(id)
		if err != nil {
			writeError(w, "Last-Event-ID must be a number!", http.StatusBadRequest)
			return
		}
	}
	// Create the stream subscriber
	sub := stream.addSubscriber(eventId, r.URL)

	go func() {
		<-r.Context().Done()

		sub.close()

		if s.autoStream && !s.autoReplay && stream.getSubscriberCount() == 0 {
			s.streamMgr.RemoveWithID(StreamID(streamID))
		}
	}()

	w.WriteHeader(http.StatusOK)
	flusher.Flush()

	// Push events to client
	for ev := range sub.connection {
		// If the data buffer is an empty string abort.
		if len(ev.Data) == 0 && len(ev.Comment) == 0 {
			break
		}
		// if the event has expired, don't send it
		if s.eventTTL != 0 && time.Now().After(ev.timestamp.Add(s.eventTTL)) {
			continue
		}
		if len(ev.Data) > 0 {
			_, _ = writeData(w, FieldId, ev.ID)

			if s.splitData {
				sd := bytes.Split(ev.Data, []byte("\n"))
				for i := range sd {
					_, _ = writeData(w, FieldData, sd[i])
				}
			} else {
				if bytes.HasPrefix(ev.Data, []byte(FieldComment)) {
					_, _ = fmt.Fprintf(w, "%s\n", ev.Data)
				} else {
					_, _ = writeData(w, FieldData, ev.Data)
				}
			}
			if len(ev.Event) > 0 {
				_, _ = writeData(w, FieldEvent, ev.Event)
			}
			if len(ev.Retry) > 0 {
				_, _ = writeData(w, FieldRetry, ev.Retry)
			}
		}
		if len(ev.Comment) > 0 {
			_, _ = writeData(w, "", ev.Comment)
		}

		_, _ = fmt.Fprint(w, "\n")

		flusher.Flush()
	}
}

func (s *Server) prepareHeaderForSSE(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	for k, v := range s.headers {
		w.Header().Set(k, v)
	}
}
