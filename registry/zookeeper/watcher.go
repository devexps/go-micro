package zookeeper

import (
	"context"
	"errors"
	"path"
	"sync/atomic"

	"github.com/devexps/go-micro/v2/registry"
	"github.com/go-zookeeper/zk"
)

var _ registry.Watcher = (*watcher)(nil)

var ErrWatcherStopped = errors.New("watcher stopped")

type watcher struct {
	ctx    context.Context
	event  chan zk.Event
	conn   *zk.Conn
	cancel context.CancelFunc

	first       uint32
	prefix      string
	serviceName string
}

func newWatcher(ctx context.Context, prefix, serviceName string, conn *zk.Conn) (*watcher, error) {
	w := &watcher{conn: conn, event: make(chan zk.Event, 1), prefix: prefix, serviceName: serviceName}
	w.ctx, w.cancel = context.WithCancel(ctx)
	go w.watch(w.ctx)
	return w, nil
}

func (w *watcher) watch(ctx context.Context) {
	for {
		// Each watch has only one validity period, so the watch is looped
		_, _, ch, err := w.conn.ChildrenW(w.prefix)
		if err != nil {
			// If the target service node has not been created
			if errors.Is(err, zk.ErrNoNode) {
				// Add watcher for the node exists
				_, _, ch, err = w.conn.ExistsW(w.prefix)
			}
			if err != nil {
				w.event <- zk.Event{Err: err}
				return
			}
		}
		select {
		case <-ctx.Done():
			return
		case ev := <-ch:
			w.event <- ev
		}
	}
}

func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
	// todo If you call next in multiple places, it may cause multi-instance information to be out of sync
	if atomic.CompareAndSwapUint32(&w.first, 0, 1) {
		return w.getServices()
	}
	select {
	case <-w.ctx.Done():
		return nil, w.ctx.Err()
	case e := <-w.event:
		if e.State == zk.StateDisconnected {
			return nil, ErrWatcherStopped
		}
		if e.Err != nil {
			return nil, e.Err
		}
		return w.getServices()
	}
}

func (w *watcher) Stop() error {
	w.cancel()
	return nil
}

func (w *watcher) getServices() ([]*registry.ServiceInstance, error) {
	servicesID, _, err := w.conn.Children(w.prefix)
	if err != nil {
		return nil, err
	}
	items := make([]*registry.ServiceInstance, 0, len(servicesID))
	for _, id := range servicesID {
		servicePath := path.Join(w.prefix, id)
		b, _, err := w.conn.Get(servicePath)
		if err != nil {
			return nil, err
		}
		item, err := unmarshal(b)
		if err != nil {
			return nil, err
		}

		// If the service name is different from watch, it will be skipped.
		if item.Name != w.serviceName {
			continue
		}

		items = append(items, item)
	}
	return items, nil
}
