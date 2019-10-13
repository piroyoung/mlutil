package linalg

type DenseVector []float32

func (v DenseVector) AsMatrix() DenseMatrix {
	m := make([][]float32, 1, len(v))
	m[0] = v
	return m
}

type DenseMatrix [][]float32
