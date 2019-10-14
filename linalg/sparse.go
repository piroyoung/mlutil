package linalg

import "gonum.org/v1/gonum/mat"

// Adapter implementations for gonum dense

type SparseVector struct {
	Values map[uint32]float32
	Length uint32
}

func (v *SparseVector) GetDense() *mat.VecDense {
	d := make([]float64, v.Length)
	for i, v := range v.Values {
		d[i] = float64(v)
	}
	return mat.NewVecDense(int(v.Length), d)
}

func (v *SparseVector) AsMatrix() SparseMatrix {
	return SparseMatrix{
		Values:    map[uint32]map[uint32]float32{0: v.Values},
		RowLength: 1,
		ColLength: v.Length,
	}
}

type SparseMatrix struct {
	Values    map[uint32]map[uint32]float32
	RowLength uint32
	ColLength uint32
}

func (m *SparseMatrix) GetDense() *mat.Dense {
	d := make([]float64, m.RowLength*m.ColLength)
	for i, row := range m.Values {
		for j, v := range row {
			d[i*m.RowLength+j] = float64(v)
		}
	}
	return mat.NewDense(int(m.RowLength), int(m.ColLength), d)
}
