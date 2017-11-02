package cmap

import "sync"

//@formatter:off

type Map struct {
    sync.RWMutex
    dict map[string]struct{}
}

func NewMap() *Map {
    return &Map{dict:make(map[string]struct{})}
}

func (m *Map) HasKey(key string) bool {
    m.RLock()
    var _, ok = m.dict[key]
    m.RUnlock()
    return ok
}

func (m *Map) Set(key string) {
    m.Lock()
    m.dict[key] = struct{}{}
    m.Unlock()
}

func (m *Map) Delete(key string) {
    m.Lock()
    delete(m.dict, key)
    m.Unlock()
}

func (m *Map) Size() int {
    m.RLock()
    var v = len(m.dict)
    m.RUnlock()
    return v
}


func (m *Map) Keys () []string {
    m.Lock()
    var keys = make([]string , m.Size())
    for k := range m.dict {
        keys = append(keys, k)
    }
    m.Unlock()
    return keys
}
