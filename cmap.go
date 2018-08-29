package cmap

import (
	"sync"
	"github.com/TopoSimplify/node"
	"sort"
)

const CacheKeySize = 6

type CacheMap struct {
	sync.RWMutex
	dict map[[CacheKeySize]int]struct{}
}

func NewCacheMap(size int) *CacheMap {
	return &CacheMap{dict: make(map[[CacheKeySize]int]struct{}, size)}
}

func (m *CacheMap) HasKey(key [CacheKeySize]int) bool {
	m.RLock()
	var _, ok = m.dict[key]
	m.RUnlock()
	return ok
}

func (m *CacheMap) Set(key [CacheKeySize]int) {
	m.Lock()
	m.dict[key] = struct{}{}
	m.Unlock()
}

func (m *CacheMap) Delete(key [CacheKeySize]int) {
	m.Lock()
	delete(m.dict, key)
	m.Unlock()
}

func (m *CacheMap) Size() int {
	m.RLock()
	var v = len(m.dict)
	m.RUnlock()
	return v
}

func (m *CacheMap) Keys() [][CacheKeySize]int {
	m.RLock()
	var keys = make([][CacheKeySize]int, m.Size())
	for k := range m.dict {
		keys = append(keys, k)
	}
	m.RUnlock()
	return keys
}

func CacheKey(a, b *node.Node) [CacheKeySize]int {
	var o = [CacheKeySize]int{
		a.Range.I, a.Range.J,
		b.Range.I, b.Range.J,
		a.Instance.Id(), b.Instance.Id(),
	}
	sort.Ints(o[:])
	return o
}

