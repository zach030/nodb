package main

import (
	"fmt"
	"nodb/mem"
)

func main() {
	s := mem.NewStorageInstance()
	s.Insert([]byte{1, 2}, []byte{2, 3})
	res, err := s.Search([]byte{1, 2})
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Println("res is:", res)
}
