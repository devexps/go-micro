package k8s

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const (
	namespace = "default"
	name      = "test"
	testKey   = "test_key"
)

var (
	keyPath    = strings.Join([]string{namespace, name, testKey}, "/")
	objectMeta = metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
		Labels: map[string]string{
			"app": "test",
		},
	}
)

func TestSource(t *testing.T) {
	s := NewSource(
		Namespace(namespace),
		LabelSelector(""),
		KubeConfig(filepath.Join(homedir.HomeDir(), ".kube", "config")),
	)
	kvs, err := s.Load()
	if err != nil {
		t.Error(err)
	}
	for _, v := range kvs {
		t.Log(v)
	}
}

func TestConfig(t *testing.T) {
	restConfig, err := rest.InClusterConfig()

	options := []Option{
		Namespace(namespace),
		LabelSelector("app=test"),
	}

	if err != nil {
		kubeConfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		restConfig, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			t.Fatal(err)
		}
		options = append(options, KubeConfig(kubeConfig))
	}
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		t.Fatal(err)
	}

	clientSetConfigMaps := clientSet.CoreV1().ConfigMaps(namespace)

	source := NewSource(options...)
	if _, err = clientSetConfigMaps.Create(context.Background(), &v1.ConfigMap{
		ObjectMeta: objectMeta,
		Data: map[string]string{
			testKey: "test config",
		},
	}, metav1.CreateOptions{}); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err = clientSetConfigMaps.Delete(context.Background(), name, metav1.DeleteOptions{}); err != nil {
			t.Error(err)
		}
	}()
	kvs, err := source.Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(kvs) != 1 || kvs[0].Key != keyPath || string(kvs[0].Value) != "test config" {
		t.Fatal("config error")
	}

	w, err := source.Watch()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = w.Stop()
	}()
	// create also produce an event, discard it
	if _, err = w.Next(); err != nil {
		t.Fatal(err)
	}

	if _, err = clientSetConfigMaps.Update(context.Background(), &v1.ConfigMap{
		ObjectMeta: objectMeta,
		Data: map[string]string{
			testKey: "new config",
		},
	}, metav1.UpdateOptions{}); err != nil {
		t.Error(err)
	}

	if kvs, err = w.Next(); err != nil {
		t.Fatal(err)
	}

	if len(kvs) != 1 || kvs[0].Key != keyPath || string(kvs[0].Value) != "new config" {
		t.Fatal("config error")
	}
}
