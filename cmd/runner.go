package main

import (
	"fmt"
	"github.com/piroyoung/scigo/hasher"
)

func main() {
	h := hasher.NewCRC32FeatureHasher(32)
	vec, err := h.GetSparseVector([]string{"apple", "orange", "banana"})
	if err != nil {
		panic(err)
	}
	fmt.Println(vec.GetDense())
}
