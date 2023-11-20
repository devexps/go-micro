package broker

import (
	"errors"
	"fmt"
	"sync"
)

type Subscriber interface {
	Options() SubscribeOptions
	Topic() string
	Unsubscribe() error
}
type SubscriberMap map[string]Subscriber

type SubscriberSyncMap struct {
	sync.RWMutex
	m SubscriberMap
}

func NewSubscriberSyncMap() *SubscriberSyncMap {
	return &SubscriberSyncMap{
		m: make(SubscriberMap),
	}
}

func (sm *SubscriberSyncMap) Add(topic string, sub Subscriber) {
	sm.Lock()
	defer sm.Unlock()

	sm.m[topic] = sub
}

func (sm *SubscriberSyncMap) Remove(topic string) error {
	sm.Lock()
	defer sm.Unlock()

	if sub, ok := sm.m[topic]; ok {
		delete(sm.m, topic)
		return sub.Unsubscribe()
	} else {
		return errors.New(fmt.Sprintf("topic[%s] not found", topic))
	}
}

func (sm *SubscriberSyncMap) Clear() {
	sm.Lock()
	defer sm.Unlock()

	for _, sub := range sm.m {
		_ = sub.Unsubscribe()
	}
	sm.m = make(SubscriberMap)
}

func (sm *SubscriberSyncMap) ForceClear() {
	sm.Lock()
	defer sm.Unlock()

	sm.m = make(SubscriberMap)
}

func (sm *SubscriberSyncMap) Get(topic string) Subscriber {
	sm.RLock()
	defer sm.RUnlock()

	return sm.m[topic]
}

func (sm *SubscriberSyncMap) Foreach(fnc func(topic string, sub Subscriber)) {
	sm.RLock()
	defer sm.RUnlock()

	for k, v := range sm.m {
		fnc(k, v)
	}
}
