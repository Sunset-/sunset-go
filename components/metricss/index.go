package metricss

import (
	"github.com/rcrowley/go-metrics"
	"sync"
)

var registryMap = map[string]metrics.Registry{}
var registryMapLock sync.RWMutex

func Registry(registry string) metrics.Registry {
	registryMapLock.RLock()
	if r, ok := registryMap[registry]; ok {
		registryMapLock.RUnlock()
		return r
	}
	registryMapLock.RUnlock()
	registryMapLock.Lock()
	r := metrics.NewRegistry()
	registryMap[registry] = r
	registryMapLock.Unlock()
	return r
}

func MonitorCount(group string, key string) metrics.Counter {
	return metrics.GetOrRegisterCounter(key, Registry(group))
}

func MonitorTimer(group string, key string) metrics.Timer {
	return metrics.GetOrRegisterTimer(key, Registry(group))
}
