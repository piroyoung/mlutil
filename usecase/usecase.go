package main

import (
	"fmt"
	"github.com/piroyoung/mlutil/hasher"
)

func main() {
	h := hasher.NewCRC32FeatureHasher(71)
	features := make([][]string, 3)
	features[0] = []string{"Are", "you", "satisfied", "with", "python2", "?"}
	features[1] = []string{"Nobody", "loves", "python2", "users"}
	features[2] = []string{"Don't", "worry", "golang", "always", "loves", "you"}
	if sparse, err := h.GetSparseMatrix(features); err == nil {
		// returns {map[0:map[12:1 23:1 24:1 59:-1 63:1 68:-1] 1:map[6:-1 42:-1 61:1 68:-1] 2:map[11:-1 14:1 18:-1 59:-1 60:1 61:1]] 3 71}
		fmt.Println(sparse)
		dense := sparse.GetDense()
		//
		fmt.Println(dense)
	} else {
		panic(err)
	}
}
