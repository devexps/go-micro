package utils

import (
	"context"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/devexps/go-micro/v2/errors"
)

type KeepAliveService struct {
	*http.Server
	lis      net.Listener
	tlsConf  *tls.Config
	endpoint *url.URL
}

// NewKeepAliveService .
func NewKeepAliveService(tlsConf *tls.Config) *KeepAliveService {
	srv := &KeepAliveService{
		tlsConf: tlsConf,
	}
	srv.Server = &http.Server{
		TLSConfig: srv.tlsConf,
	}
	return srv
}

// Endpoint return a real address to registry endpoint.
func (s *KeepAliveService) Endpoint() (*url.URL, error) {
	if err := s.generateEndpoint(); err != nil {
		return nil, err
	}
	return s.endpoint, nil
}

// Start starts the HTTP server.
func (s *KeepAliveService) Start() error {
	if err := s.generateEndpoint(); err != nil {
		return err
	}
	var err error
	if s.tlsConf != nil {
		err = s.ServeTLS(s.lis, "", "")
	} else {
		err = s.Serve(s.lis)
	}
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stops the HTTP server.
func (s *KeepAliveService) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}

func (s *KeepAliveService) generatePort(min, max int) int {
	return rand.Intn(max-min) + min
}

func (s *KeepAliveService) generateEndpoint() error {
	if s.endpoint != nil {
		return nil
	}
	for {
		port := s.generatePort(10000, 65535)
		host := ""
		if itf, ok := os.LookupEnv("GOMICRO_TRANSPORT_KEEPALIVE_INTERFACE"); ok {
			h, err := getIPAddress(itf)
			if err != nil {
				return err
			}
			host = h
		}
		addr := fmt.Sprintf("%s:%d", host, port)
		lis, err := net.Listen("tcp", addr)
		if err == nil && lis != nil {
			s.lis = lis
			endpoint, _ := url.Parse("tcp://" + addr)
			s.endpoint = endpoint
			return nil
		}
	}
}

func getIPAddress(interfaceName string) (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range interfaces {
		if iface.Name == interfaceName {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}
			// Get the first IPv4 address
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					return ipnet.IP.String(), nil
				}
			}
			return "", fmt.Errorf("no IPv4 address found for interface %s", interfaceName)
		}
	}
	return "", fmt.Errorf("interface %s not found", interfaceName)
}
