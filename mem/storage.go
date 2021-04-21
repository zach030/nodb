package mem

import (
	"errors"
	"sync"
)

type Atom []byte

func (a Atom) ToString() string {
	return string(a)
}

var (
	KeyNotFound = errors.New("not found key in nodb")
)

// 内存数据库
// 语法分析 + 索引结构
type Storage struct {
	lock sync.Mutex
	DB   map[string]Atom
}

func NewStorageInstance() *Storage {
	return &Storage{
		DB: make(map[string]Atom, 0),
	}
}

func (s *Storage) Insert(key, value Atom) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.DB[key.ToString()] = value
}

func (s *Storage) Search(key Atom) (Atom, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if v, ok := s.DB[key.ToString()]; ok {
		return v, nil
	}
	return nil, KeyNotFound
}
