package k8s

import (
	"context"
	"time"

	jsoniter "github.com/json-iterator/go"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	listerv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"

	"github.com/devexps/go-micro/v2/registry"
)

const (
	// LabelsKeyServiceID is used to define the ID of the service
	LabelsKeyServiceID = "gomicro-service-id"

	// LabelsKeyServiceName is used to define the name of the service
	LabelsKeyServiceName = "gomicro-service-app"

	// LabelsKeyServiceVersion is used to define the version of the service
	LabelsKeyServiceVersion = "gomicro-service-version"

	// AnnotationsKeyMetadata is used to define the metadata of the service
	AnnotationsKeyMetadata = "gomicro-service-metadata"

	// AnnotationsKeyProtocolMap is used to define the protocols of the service
	// Through the value of this field, GoMicro can obtain the application layer protocol corresponding to the port
	// Example value: {"8080": "http", "9090": "grpc"}
	AnnotationsKeyProtocolMap = "gomicro-service-protocols"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

// The Registry simply implements service discovery based on Kubernetes
type Registry struct {
	clientSet       *kubernetes.Clientset
	informerFactory informers.SharedInformerFactory
	podInformer     cache.SharedIndexInformer
	podLister       listerv1.PodLister

	stopCh chan struct{}
}

// NewRegistry is used to initialize the Registry
func NewRegistry(clientSet *kubernetes.Clientset) *Registry {
	informerFactory := informers.NewSharedInformerFactory(clientSet, time.Minute*10)
	podInformer := informerFactory.Core().V1().Pods().Informer()
	podLister := informerFactory.Core().V1().Pods().Lister()
	return &Registry{
		clientSet:       clientSet,
		informerFactory: informerFactory,
		podInformer:     podInformer,
		podLister:       podLister,
		stopCh:          make(chan struct{}),
	}
}

// Register is used to register services
func (s *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	// Generate Metadata
	metadataVal, err := marshal(service.Metadata)
	if err != nil {
		return err
	}

	// Generate ProtocolMap
	protocolMap, err := getProtocolMapByEndpoints(service.Endpoints)
	if err != nil {
		return err
	}
	protocolMapVal, err := marshal(protocolMap)
	if err != nil {
		return err
	}

	// Generate Patch
	patchBytes, err := jsoniter.Marshal(map[string]interface{}{
		"metadata": metav1.ObjectMeta{
			Labels: map[string]string{
				LabelsKeyServiceID:      service.ID,
				LabelsKeyServiceName:    service.Name,
				LabelsKeyServiceVersion: service.Version,
			},
			Annotations: map[string]string{
				AnnotationsKeyMetadata:    metadataVal,
				AnnotationsKeyProtocolMap: protocolMapVal,
			},
		},
	})
	if err != nil {
		return err
	}

	// applies the patch and returns the patched pod
	if _, err = s.clientSet.
		CoreV1().
		Pods(GetNamespace()).
		Patch(ctx, GetPodName(), types.StrategicMergePatchType, patchBytes, metav1.PatchOptions{}); err != nil {
		return err
	}
	return nil
}

// Deregister the registration.
func (s *Registry) Deregister(ctx context.Context, _ *registry.ServiceInstance) error {
	return s.Register(ctx, &registry.ServiceInstance{
		Metadata: map[string]string{},
	})
}

// GetService return the service instances in memory according to the service name.
func (s *Registry) GetService(ctx context.Context, name string) ([]*registry.ServiceInstance, error) {
	pods, err := s.podLister.List(labels.SelectorFromSet(map[string]string{
		LabelsKeyServiceName: name,
	}))
	if err != nil {
		return nil, err
	}
	ret := make([]*registry.ServiceInstance, 0, len(pods))
	for _, pod := range pods {
		if pod.Status.Phase != corev1.PodRunning {
			continue
		}
		instance, err := getServiceInstanceFromPod(pod)
		if err != nil {
			return nil, err
		}
		ret = append(ret, instance)
	}
	return ret, nil
}

// Watch creates a watcher according to the service name.
func (s *Registry) Watch(ctx context.Context, name string) (registry.Watcher, error) {
	stopCh := make(chan struct{}, 1)
	announcement := make(chan []*registry.ServiceInstance, 1)

	s.podInformer.AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: func(obj interface{}) bool {
			select {
			case <-stopCh:
				return false
			case <-s.stopCh:
				return false
			default:
				pod := obj.(*corev1.Pod)
				val := pod.GetLabels()[LabelsKeyServiceName]
				return val == name
			}
		},
		Handler: cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				s.sendLatestInstances(ctx, name, announcement)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				s.sendLatestInstances(ctx, name, announcement)
			},
			DeleteFunc: func(obj interface{}) {
				s.sendLatestInstances(ctx, name, announcement)
			},
		},
	})
	return NewIterator(announcement, stopCh), nil
}

// Start is used to start the Registry
// It is non-blocking
func (s *Registry) Start() {
	s.informerFactory.Start(s.stopCh)
	if !cache.WaitForCacheSync(s.stopCh, s.podInformer.HasSynced) {
		return
	}
}

// Close is used to close the Registry
// After closing, any callbacks generated by Watch will not be executed
func (s *Registry) Close() {
	select {
	case <-s.stopCh:
	default:
		close(s.stopCh)
	}
}

func (s *Registry) sendLatestInstances(ctx context.Context, name string, announcement chan []*registry.ServiceInstance) {
	instances, err := s.GetService(ctx, name)
	if err != nil {
		// something went wrong
		panic(err)
	}
	announcement <- instances
}
