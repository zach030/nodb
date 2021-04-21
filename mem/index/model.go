package index

import "fmt"

type BPlusTree struct {
}

func NewBPlusTree() *BPlusTree {
	return &BPlusTree{}
}

func (B BPlusTree) InsertElement(s, atom Atom) {
	panic("implement me")
}

func (B BPlusTree) GetElement(s Atom) Atom {
	panic("implement me")
}

func (B BPlusTree) RemoveElement(s Atom) {
	panic("implement me")
}

type HashIndex struct {
}

func NewHashIndex() *HashIndex {
	return &HashIndex{}
}

func (h HashIndex) InsertElement(s, atom Atom) {
	panic("implement me")
}

func (h HashIndex) GetElement(s Atom) Atom {
	panic("implement me")
}

func (h HashIndex) RemoveElement(s Atom) {
	panic("implement me")
}

type LSMIndex struct {
}

func NewLSMIndex() *LSMIndex {
	return &LSMIndex{}
}

func (L LSMIndex) InsertElement(s, atom Atom) {
	panic("implement me")
}

func (L LSMIndex) GetElement(s Atom) Atom{
	panic("implement me")
}

func (L LSMIndex) RemoveElement(s Atom) {
	panic("implement me")
}

type ArrayIndex struct {
}

func NewArrayIndex() *ArrayIndex {
	return &ArrayIndex{}
}

func (a ArrayIndex) InsertElement(s, atom Atom) {
	panic("implement me")
}

func (a ArrayIndex) GetElement(s Atom) Atom {
	panic("implement me")
}

func (a ArrayIndex) RemoveElement(s Atom) {
	panic("implement me")
}

type MapIndex struct {
	DB map[string]Atom
}

func NewMapIndex() *MapIndex {
	return &MapIndex{
		DB: make(map[string]Atom, 0),
	}
}

func (m *MapIndex) InsertElement(s, atom Atom) {
	m.DB[s.ToString()] = atom
	fmt.Printf("map index success insert element\n")
}

func (m *MapIndex) GetElement(s Atom) Atom {
	if v, ok := m.DB[s.ToString()]; ok {
		fmt.Printf("map index success get value:%v\n",v)
		return v
	}
	return nil
}

func (m *MapIndex) RemoveElement(s Atom) {
	delete(m.DB, s.ToString())
}
