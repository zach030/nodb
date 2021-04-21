package index

import (
	"errors"
	"fmt"
)

type StorageIndex interface {
	InsertElement(Atom, Atom)
	GetElement(Atom)Atom
	RemoveElement(Atom)
}

var (
	NoIndexModel = errors.New("not found specific index model in nodb")
)

const (
	Map   = "map"
	BPlus = "b-plus"
	Hash  = "hash"
	LSM   = "lsm"
	Array = "array"
)

type Engine struct {
	IndexPool map[string]StorageIndex
}

func NewIndexEngine(index string) *Engine {
	var e = &Engine{IndexPool: make(map[string]StorageIndex)}
	err := e.initIndex(index)
	if err != nil {
		return nil
	}
	return e
}

func (e *Engine) initIndex(index string) error {
	model, err := e.GetIndexModel(index)
	if err != nil {
		return err
	}
	fmt.Printf("index model is :%+v\n", model)
	return nil
}

func (e *Engine) GetIndexModel(name string) (StorageIndex, error) {
	if model, ok := e.IndexPool[name]; ok {
		return model, nil
	}
	switch name {
	case BPlus:
		bpModel := NewBPlusTree()
		e.IndexPool[name] = bpModel
		return bpModel, nil
	case Hash:
		hashModel := NewHashIndex()
		e.IndexPool[name] = hashModel
		return hashModel, nil
	case LSM:
		lsmModel := NewLSMIndex()
		e.IndexPool[name] = lsmModel
		return lsmModel, nil
	case Array:
		arrModel := NewArrayIndex()
		e.IndexPool[name] = arrModel
		return arrModel, nil
	case Map:
		mapModel := NewMapIndex()
		e.IndexPool[name] = mapModel
		return mapModel, nil
	default:
		return nil, NoIndexModel
	}
}

func (e *Engine) Insert(index string, key, value Atom) error {
	model, err := e.GetIndexModel(index)
	if err != nil {
		return err
	}
	fmt.Printf("current index model is:%+v\n",model)
	model.InsertElement(key, value)
	return nil
}

func (e *Engine) Search(index string, key Atom) (Atom, error) {
	model, err := e.GetIndexModel(index)
	if err != nil {
		return nil, err
	}
	res := model.GetElement(key)
	return res, nil
}
