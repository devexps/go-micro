package zookeeper

import (
	"encoding/json"

	"github.com/devexps/go-micro/v2/registry"
)

func marshal(si *registry.ServiceInstance) ([]byte, error) {
	return json.Marshal(si)
}

func unmarshal(data []byte) (si *registry.ServiceInstance, err error) {
	err = json.Unmarshal(data, &si)
	return
}
