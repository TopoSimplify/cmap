package cmap

import (
    "sync"
)

type Map struct {
    sync.RWMutex
    dict map[string]bool
}

func (m *Map) Get(key string) bool {
    m.RLock()
    var v = m.dict[key]
    m.RUnlock()
    return v
}

func (m *Map) Set(key string, val bool) {
    m.Lock()
    m.dict[key] = val
    m.Unlock()
}

func (m *Map) Size() int {
    m.RLock()
    var v = len(m.dict)
    m.RUnlock()
    return v
}
