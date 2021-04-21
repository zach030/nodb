package mem

import (
	"errors"
	"nodb/mem/index"
	"sync"
)


var (
	KeyNotFound = errors.New("not found key in nodb")
)

type Option struct {
	Index string
}

var DefaultOption = &Option{
	Index: index.Map,
}

// 内存数据库
// 语法分析 + 索引结构
type Storage struct {
	CurrentIndex string
	lock         sync.Mutex
	Engine       *index.Engine
}

func NewStorageInstance(opt *Option) *Storage {
	if opt == nil {
		opt = DefaultOption
	}
	return &Storage{
		CurrentIndex: opt.Index,
		Engine:       index.NewIndexEngine(opt.Index),
	}
}

func (s *Storage) Insert(key, value index.Atom) {
	s.lock.Lock()
	defer s.lock.Unlock()
	err := s.Engine.Insert(s.CurrentIndex, key, value)
	if err != nil {

	}
}

func (s *Storage) Search(key index.Atom) (index.Atom, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Engine.Search(s.CurrentIndex, key)
}
