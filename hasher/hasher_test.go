package hasher

import (
	"testing"
)

func TestCrc32FeatureHasher_GetHashedFeature(t *testing.T) {
	hasher := NewCRC32FeatureHasher(1024)
	feature, err := hasher.GetHashedFeature("example")
	if err != nil {
		t.Error()
	} else {
		t.Log(feature)
	}
}

func TestCrc32FeatureHasher_GetSparseVector(t *testing.T) {
	h := NewCRC32FeatureHasher(32)
	features := []string{"apple", "orange", "banana", "pine",}
	vec, err := h.GetSparseVector(features)
	if err != nil {
		t.Error()
	}
	m := vec.AsMatrix()
	t.Log(m.GetDense())
}
