package bitset

import (
	bse "go-bloom-axon/bitset_error"
)

type BitSet struct {
	size uint64
	bitMap []uint64
}

func New(size uint64) *BitSet {

	return &BitSet{
		size : size ,
		bitMap : make([]uint64 , (size + 63)/64),
	}
}

func (bs *BitSet) Set(n uint64) error {
	if n >= bs.size {
		return bse.ErrIndexOutOfBounds
	}

	targetKey := n/64
	targetBit := n%64

	bs.bitMap[targetKey] |= (1<<targetBit)
	return nil
}

func (bs *BitSet) Size() uint64 {
	return bs.size
}

func (bs *BitSet) Get(n uint64) (bool, error) {
	if n >= bs.size {
		return false, bse.ErrIndexOutOfBounds
	}

	targetKey := n/64
	targetBit := n%64

	return (bs.bitMap[targetKey] & (1<<targetBit)) != 0, nil
}
