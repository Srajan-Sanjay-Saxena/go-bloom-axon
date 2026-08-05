package counterbloomfilter

import (
	"errors"

	"github.com/twmb/murmur3"
)

var (
	ErrUnderflow     = errors.New("counter underflow: item may not have been added")
	ErrCounterMaxed  = errors.New("counter overflow: max count of 255 reached for this slot")
)

// CountingBloomFilter uses a []uint8 counter array instead of a bitset.
// Each slot holds a count 0–255. Add increments, Remove decrements.
// This makes deletion safe — something a standard Bloom Filter cannot do.
type CountingBloomFilter struct {
	counters []uint8
	size     uint64
	hashCnt  uint64
}

func New(size uint64) *CountingBloomFilter {
	counterSize := size * 10
	return &CountingBloomFilter{
		counters: make([]uint8, counterSize),
		size:     counterSize,
		hashCnt:  7,
	}
}

func (cbf *CountingBloomFilter) hashes(item []byte) (uint64, uint64) {
	return murmur3.Sum128(item)
}

func (cbf *CountingBloomFilter) positions(item []byte) []uint64 {
	h1, h2 := cbf.hashes(item)
	pos := make([]uint64, cbf.hashCnt)
	for i := uint64(0); i < cbf.hashCnt; i++ {
		pos[i] = (h1 + i*h2) % cbf.size
	}
	return pos
}

// Add increments the counter at each of the k hash positions.
// Returns ErrCounterMaxed if any slot is already at 255.
func (cbf *CountingBloomFilter) Add(item []byte) error {
	for _, p := range cbf.positions(item) {
		if cbf.counters[p] == 255 {
			return ErrCounterMaxed
		}
	}
	for _, p := range cbf.positions(item) {
		cbf.counters[p]++
	}
	return nil
}

// Remove decrements the counter at each of the k hash positions.
// Returns ErrUnderflow if any counter is already 0 — item was never added.
func (cbf *CountingBloomFilter) Remove(item []byte) error {
	for _, p := range cbf.positions(item) {
		if cbf.counters[p] == 0 {
			return ErrUnderflow
		}
	}
	for _, p := range cbf.positions(item) {
		cbf.counters[p]--
	}
	return nil
}

// Contains returns true if all counters at the k positions are > 0.
// False means definitely not present. True means probably present.
func (cbf *CountingBloomFilter) Contains(item []byte) bool {
	for _, p := range cbf.positions(item) {
		if cbf.counters[p] == 0 {
			return false
		}
	}
	return true
}
