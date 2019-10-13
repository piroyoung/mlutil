package hasher

import (
	"encoding/binary"
	"github.com/piroyoung/scigo/linalg"
	"hash/crc32"
)

type HashedFeature struct {
	Index uint32
	Value float32
}

type FeatureHasher interface {
	GetHashedFeature(feature string) (HashedFeature, error)
	GetSparseVector(features []string) (linalg.SparseVector, error)
}

type crc32FeatureHasher struct {
	modular uint32
}

func NewCRC32FeatureHasher(modular uint32) FeatureHasher {
	return &crc32FeatureHasher{
		modular: modular,
	}
}

func (h *crc32FeatureHasher) getCheckDigit(value uint32) float32 {
	bytes := make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(bytes, value)
	switch crc32.ChecksumIEEE(bytes) % 2 {
	case 0:
		return -1.0
	default:
		return 1.0
	}
}

func (h *crc32FeatureHasher) GetHashedFeature(feature string) (HashedFeature, error) {
	sum := crc32.ChecksumIEEE([]byte(feature))
	return HashedFeature{
		Index: sum % h.modular,
		Value: h.getCheckDigit(sum),
	}, nil
}

func (h *crc32FeatureHasher) GetSparseVector(features []string) (linalg.SparseVector, error) {
	values := map[uint32]float32{}
	for _, value := range features {
		hashed, err := h.GetHashedFeature(value)
		if err != nil {
			return linalg.SparseVector{}, err
		}
		values[hashed.Index] += hashed.Value
	}
	return linalg.SparseVector{
		Values: values,
		Length: h.modular,
	}, nil
}
