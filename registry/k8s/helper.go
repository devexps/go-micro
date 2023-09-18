package k8s

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"

	corev1 "k8s.io/api/core/v1"

	"github.com/devexps/go-micro/v2/registry"
)

////////////// ProtocolMap ////////////
type protocolMap map[string]string

func (m protocolMap) GetProtocol(port int32) string {
	return m[strconv.Itoa(int(port))]
}

// //////////// K8S Runtime ////////////

// ServiceAccountNamespacePath defines the location of the namespace file
const ServiceAccountNamespacePath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"

var currentNamespace = LoadNamespace()

// LoadNamespace is used to get the current namespace from the file
func LoadNamespace() string {
	data, err := os.ReadFile(ServiceAccountNamespacePath)
	if err != nil {
		return ""
	}
	return string(data)
}

// GetNamespace is used to get the namespace of the Pod where the current container is located
func GetNamespace() string {
	return currentNamespace
}

// GetPodName is used to get the name of the Pod where the current container is located
func GetPodName() string {
	return os.Getenv("HOSTNAME")
}

// //////////// Error Definition ////////////

// ErrIteratorClosed defines the error that the iterator is closed
var ErrIteratorClosed = errors.New("iterator closed")

// ErrorHandleResource defines the error that cannot handle K8S resources normally
type ErrorHandleResource struct {
	Namespace string
	Name      string
	Reason    error
}

// Error implements the error interface
func (err *ErrorHandleResource) Error() string {
	return fmt.Sprintf("failed to handle resource(namespace=%s, name=%s): %s",
		err.Namespace, err.Name, err.Reason)
}

// //////////// Iterator ////////////

// Iterator performs the conversion from channel to iterator
// It reads the latest changes from the `chan []*registry.ServiceInstance`
// And the outside can sense the closure of Iterator through stopCh
type Iterator struct {
	ch     chan []*registry.ServiceInstance
	stopCh chan struct{}
}

// NewIterator is used to initialize Iterator
func NewIterator(channel chan []*registry.ServiceInstance, stopCh chan struct{}) *Iterator {
	return &Iterator{
		ch:     channel,
		stopCh: stopCh,
	}
}

// Next will block until ServiceInstance changes
func (iter *Iterator) Next() ([]*registry.ServiceInstance, error) {
	select {
	case instances := <-iter.ch:
		return instances, nil
	case <-iter.stopCh:
		return nil, ErrIteratorClosed
	}
}

// Stop is used to close the iterator
func (iter *Iterator) Stop() error {
	select {
	case <-iter.stopCh:
	default:
		close(iter.stopCh)
	}
	return nil
}

//////////////////////////////////////////

func marshal(in interface{}) (string, error) {
	return jsoniter.MarshalToString(in)
}

func unmarshal(data string, in interface{}) error {
	return jsoniter.UnmarshalFromString(data, in)
}

func isEmptyObjectString(s string) bool {
	switch s {
	case "", "{}", "null", "nil", "[]":
		return true
	}
	return false
}

func getProtocolMapByEndpoints(endpoints []string) (protocolMap, error) {
	ret := protocolMap{}
	for _, endpoint := range endpoints {
		u, err := url.Parse(endpoint)
		if err != nil {
			return nil, err
		}
		ret[u.Port()] = u.Scheme
	}
	return ret, nil
}

func getMetadataFromPod(pod *corev1.Pod) (map[string]string, error) {
	metadata := map[string]string{}
	if s := pod.Annotations[AnnotationsKeyMetadata]; !isEmptyObjectString(s) {
		err := unmarshal(s, &metadata)
		if err != nil {
			return nil, &ErrorHandleResource{Namespace: pod.Namespace, Name: pod.Name, Reason: err}
		}
	}
	return metadata, nil
}

func getProtocolMapFromPod(pod *corev1.Pod) (protocolMap, error) {
	protoMap := protocolMap{}
	if s := pod.Annotations[AnnotationsKeyProtocolMap]; !isEmptyObjectString(s) {
		err := unmarshal(s, &protoMap)
		if err != nil {
			return nil, &ErrorHandleResource{Namespace: pod.Namespace, Name: pod.Name, Reason: err}
		}
	}
	return protoMap, nil
}

func getServiceInstanceFromPod(pod *corev1.Pod) (*registry.ServiceInstance, error) {
	podIP := pod.Status.PodIP
	podLabels := pod.GetLabels()
	// Get Metadata
	metadata, err := getMetadataFromPod(pod)
	if err != nil {
		return nil, err
	}
	// Get Protocols Definition
	protocolMap, err := getProtocolMapFromPod(pod)
	if err != nil {
		return nil, err
	}
	// Get Endpoints
	var endpoints []string
	for _, container := range pod.Spec.Containers {
		for _, cp := range container.Ports {
			port := cp.ContainerPort
			protocol := protocolMap.GetProtocol(port)
			if protocol == "" {
				if cp.Name != "" {
					protocol = strings.Split(cp.Name, "-")[0]
				} else {
					protocol = string(cp.Protocol)
				}
			}
			addr := fmt.Sprintf("%s://%s:%d", protocol, podIP, port)
			endpoints = append(endpoints, addr)
		}
	}
	return &registry.ServiceInstance{
		ID:        podLabels[LabelsKeyServiceID],
		Name:      podLabels[LabelsKeyServiceName],
		Version:   podLabels[LabelsKeyServiceVersion],
		Metadata:  metadata,
		Endpoints: endpoints,
	}, nil
}
