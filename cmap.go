package cmap
//@formatter:off

import (
    "sync"
)

type Map struct {
    sync.RWMutex
    dict map[string]bool
}

func (m *Map) Get(key string) bool {
    m.RLock(); defer m.RUnlock()
    var v = m.dict[key]
    return v
}

func (m *Map) Set(key string, val bool) {
    m.Lock(); defer m.Unlock()
    m.dict[key] = val
}

func (m *Map) Size() int {
    m.RLock(); defer m.RUnlock()
    var v = len(m.dict)
    return v
}

func (m *Map) Iter (fn func (k string , v bool))  {
    m.Lock(); defer m.Unlock()
    for k, v := range m.dict {fn(k, v)}
}
