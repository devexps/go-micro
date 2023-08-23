package sse

import (
	"fmt"
	"net/http"
)

const (
	FieldId      = "id"
	FieldData    = "data"
	FieldEvent   = "event"
	FieldRetry   = "retry"
	FieldComment = ":"
)

func writeData(w http.ResponseWriter, field string, value []byte) (int, error) {
	return fmt.Fprintf(w, "%s: %s\n", field, value)
}

func writeError(w http.ResponseWriter, message string, status int) {
	http.Error(w, message, status)
}
