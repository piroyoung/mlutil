# mlutil

My personal ml utils


# modules 
## Feature Hashing
```go
package main

import (
	"fmt"
	"github.com/piroyoung/mlutil/hasher"
	"strings"
)

func main() {
	// usecase for feature hashing
	h := hasher.NewCRC32FeatureHasher(1023)
	features := make([][]string, 3)
	features[0] = strings.Split("Are you satisfied with python2 ?", " ")
	features[1] = strings.Split("Nobody loves python2 users", " ")
	features[2] = strings.Split("Don't worry golang always loves you", " ")
	if sparse, err := h.GetSparseMatrix(features); err == nil {
		// returns {map[0:map[348:-1 353:1 479:-1 552:1 675:1 977:1] 1:map[27:-1 160:-1 210:1 479:-1] 2:map[49:1 210:1 348:-1 810:-1 918:-1 928:1]] 3 1023}
		fmt.Println(sparse)

		// Gonum form dense matrix
		dense := sparse.GetDense()
		fmt.Println(dense)

	} else {
		panic(err)
	}
}

```