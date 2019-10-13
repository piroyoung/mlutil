package linalg

type SparseVector struct {
	Values map[uint32]float32
	Length uint32
}

func (v *SparseVector) GetDense() DenseVector {
	d := make([]float32, v.Length)
	for i, v := range v.Values {
		d[i] = v
	}
	return d
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

func (m *SparseMatrix) GetDense() DenseMatrix {
	d := make([][]float32, m.RowLength, m.ColLength)
	for i, row := range m.Values {
		for j, value := range row {
			d[i][j] = value
		}
	}
	return d
}
