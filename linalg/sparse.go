package linalg

import "gonum.org/v1/gonum/mat"

// Adapter implementations for gonum dense

type SparseVector struct {
	Values map[uint32]float32
	N      uint32
}

func (v *SparseVector) GetDense() *mat.VecDense {
	d := make([]float64, v.N)
	for i, v := range v.Values {
		d[i] = float64(v)
	}
	return mat.NewVecDense(int(v.N), d)
}

func (v *SparseVector) AsMatrix() SparseMatrix {
	return SparseMatrix{
		Values: map[uint32]map[uint32]float32{0: v.Values},
		N:      1,
		M:      v.N,
	}
}

type SparseMatrix struct {
	Values map[uint32]map[uint32]float32
	N      uint32
	M      uint32
}

func (m *SparseMatrix) GetDense() *mat.Dense {
	d := make([]float64, m.N*m.M)
	for i, row := range m.Values {
		for j, v := range row {
			d[i*m.M+j] = float64(v)
		}
	}
	return mat.NewDense(int(m.N), int(m.M), d)
}
