// This file has automatically been generated on Wed Feb 26 15:50:52 +05 2020.
// DO NOT EDIT.
package sync

import (
	"sync"
	_ "unsafe"
)

//go:linkname maploadorstore sync.sub_maploadorstore
func maploadorstore(m *sync.Map, key, value interface{}) (interface{}, bool) {
	return m.LoadOrStore(key, value)
}

//go:linkname MapLoadOrStore sync.sub_maploadorstore
//go:noescape
func MapLoadOrStore(m *sync.Map, key, value interface{}) (interface{}, bool)

//go:linkname maprange sync.sub_maprange
func maprange(m *sync.Map, f func(key, value interface{}) bool) {
	m.Range(f)
}

//go:linkname MapRange sync.sub_maprange
//go:noescape
func MapRange(m *sync.Map, f func(key, value interface{}) bool)

//go:linkname poolget sync.sub_poolget
func poolget(p *sync.Pool,) interface{} {
	return p.Get()
}

//go:linkname PoolGet sync.sub_poolget
//go:noescape
func PoolGet(p *sync.Pool,) interface{}

//go:linkname rwmutexrlocker sync.sub_rwmutexrlocker
func rwmutexrlocker(rw *sync.RWMutex,) sync.Locker {
	return rw.RLocker()
}

//go:linkname RWMutexRLocker sync.sub_rwmutexrlocker
//go:noescape
func RWMutexRLocker(rw *sync.RWMutex,) sync.Locker

//go:linkname NewCond sync.NewCond
//go:noescape
func NewCond(l sync.Locker,) *sync.Cond

//go:linkname mapload sync.sub_mapload
func mapload(m *sync.Map, key interface{}) (interface{}, bool) {
	return m.Load(key)
}

//go:linkname MapLoad sync.sub_mapload
//go:noescape
func MapLoad(m *sync.Map, key interface{}) (interface{}, bool)
