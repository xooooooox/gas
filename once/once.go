// used for business procedures that are executed only once

package once

import (
	"sync"
)

// default, map of box value type is interface{}

type Int64 struct {
	Mutex *sync.Mutex
	Box   map[int64]interface{}
}

func NewInt64() *Int64 {
	return &Int64{
		Mutex: &sync.Mutex{},
		Box:   map[int64]interface{}{},
	}
}

func (s *Int64) Put(key int64, value interface{}) bool {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok := s.Box[key]
	if ok {
		return false
	}
	s.Box[key] = value
	return true
}

func (s *Int64) Del(key int64) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok := s.Box[key]
	if ok {
		delete(s.Box, key)
	}
}

func (s *Int64) Has(key int64) (ok bool) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok = s.Box[key]
	return
}

func (s *Int64) Get(key int64) (value interface{}) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	val, ok := s.Box[key]
	if ok {
		value = val
	}
	return
}

type String struct {
	Mutex *sync.Mutex
	Box   map[string]interface{}
}

func NewString() *String {
	return &String{
		Mutex: &sync.Mutex{},
		Box:   map[string]interface{}{},
	}
}

func (s *String) Put(key string, value interface{}) bool {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok := s.Box[key]
	if ok {
		return false
	}
	s.Box[key] = value
	return true
}

func (s *String) Del(key string) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok := s.Box[key]
	if ok {
		delete(s.Box, key)
	}
}

func (s *String) Has(key string) (ok bool) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok = s.Box[key]
	return
}

func (s *String) Get(key string) (value interface{}) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	val, ok := s.Box[key]
	if ok {
		value = val
	}
	return
}

// map of box value type is bool

type Int64Bool struct {
	Mutex *sync.Mutex
	Box   map[int64]bool
}

func NewInt64Bool() *Int64Bool {
	return &Int64Bool{
		Mutex: &sync.Mutex{},
		Box:   map[int64]bool{},
	}
}

func (s *Int64Bool) Put(key int64, value bool) bool {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok := s.Box[key]
	if ok {
		return false
	}
	s.Box[key] = value
	return true
}

func (s *Int64Bool) Del(key int64) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok := s.Box[key]
	if ok {
		delete(s.Box, key)
	}
}

func (s *Int64Bool) Has(key int64) (ok bool) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok = s.Box[key]
	return
}

func (s *Int64Bool) Get(key int64) (value bool) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	val, ok := s.Box[key]
	if ok {
		value = val
	}
	return
}

type StringBool struct {
	Mutex *sync.Mutex
	Box   map[string]bool
}

func NewStringBool() *String {
	return &String{
		Mutex: &sync.Mutex{},
		Box:   map[string]interface{}{},
	}
}

func (s *StringBool) Put(key string, value bool) bool {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok := s.Box[key]
	if ok {
		return false
	}
	s.Box[key] = value
	return true
}

func (s *StringBool) Del(key string) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok := s.Box[key]
	if ok {
		delete(s.Box, key)
	}
}

func (s *StringBool) Has(key string) (ok bool) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, ok = s.Box[key]
	return
}

func (s *StringBool) Get(key string) (value bool) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	val, ok := s.Box[key]
	if ok {
		value = val
	}
	return
}
