package kafka

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testBrokers = "localhost:9092"
	testTopic   = "logger.sensor.ts"
	testGroupId = "fx-group"
)

func TestCreateTopic(t *testing.T) {
	err := CreateTopic(testBrokers, testTopic, 1, 1)
	assert.Nil(t, err)
}

func TestDeleteTopic(t *testing.T) {
	err := DeleteTopic(testBrokers, testTopic)
	assert.Nil(t, err)
}
