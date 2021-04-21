package index

type Atom []byte

func (a Atom) ToString() string {
	return string(a)
}
