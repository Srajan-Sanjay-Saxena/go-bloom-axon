package bloomfilter

import (
	bs "go-bloom-axon/bitset"
	"github.com/twmb/murmur3"
)

type BloomFilter struct {
	bitset *bs.BitSet
	hashCnt uint32
}

func New(size uint64) *BloomFilter {
	
	var filterSize uint64 = size * 10
	var hashCount uint32 = 7
	return &BloomFilter{
		bitset: bs.New(filterSize),
		hashCnt: hashCount,
	}
}

func (bf *BloomFilter) hashes(item []byte) (uint64, uint64) {
	return murmur3.Sum128(item)
}

func (bf *BloomFilter) Add(item []byte) {
	h1, h2 := bf.hashes(item)
	size := bf.bitset.Size()
	for i := uint64(0); i < uint64(bf.hashCnt); i++ {
		// we can simulate multiple hashes from pair of hashes
		bf.bitset.Set((h1 + i*h2) % size)
	}
}

func (bf *BloomFilter) Contains(item []byte) bool {
	h1, h2 := bf.hashes(item)
	size := bf.bitset.Size()
	for i := uint64(0); i < uint64(bf.hashCnt); i++ {
		if ok, _ := bf.bitset.Get((h1 + i*h2) % size); !ok {
			return false
		}
	}
	return true
}
