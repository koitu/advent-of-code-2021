package main

import (
	"strconv"
	"strings"
)

type bitSet []bool

func (b bitSet) toInt64() int64 {
	n := int64(0)
	l := len(b)

	if len(b) > 62 {
		panic("may be attempting to create value greater than max value for int64")
	}

	for i := 1; i <= l; i++ {
		if b[i-1] {
			n += 1 << uint(l-i)
		}
	}

	return n
}

func (b *bitSet) append4Bits(n uint8) {
	new := bitSet{}
	for i := 0; i < 4; i++ {
		new = append(new, !((1<<uint(3-i))&n == 0))
	}
	(*b) = append((*b), new...)
}

func (b *bitSet) pop(n int64) bitSet {
	r := (*b)[:n]
	*b = (*b)[n:]
	return r
}

func (b *bitSet) popInt64(n int64) int64 {
	return b.pop(n).toInt64()
}

func (b *bitSet) handlePacket() (versionTotal, expVal int64) {
	versionTotal = b.popInt64(3)
	pktType := b.popInt64(3)

	switch pktType {
	case 4:
		valueBits := bitSet{}
		for b.popInt64(1) == 1 {
			valueBits = append(valueBits, b.pop(4)...)
		}
		valueBits = append(valueBits, b.pop(4)...)

		return versionTotal, valueBits.toInt64()

	default:
		lengthType := b.popInt64(1)
		expValues := []int64{}

		switch lengthType {
		case 1:
			subPkts := b.popInt64(11)

			for i := 0; i < int(subPkts); i++ {
				ver, val := b.handlePacket()
				versionTotal += ver
				expValues = append(expValues, val)
			}
		case 0:
			subPktLen := b.popInt64(15)

			subBitSet := make(bitSet, subPktLen)
			copy(subBitSet, (*b)[:subPktLen])
			b.pop(subPktLen)

			for len(subBitSet) > 0 {
				ver, val := subBitSet.handlePacket()
				versionTotal += ver
				expValues = append(expValues, val)
			}
		}

		expVal = 0
		switch pktType {
		case 0:
			for _, v := range expValues {
				expVal += v
			}

		case 1:
			expVal = expValues[0]
			for i := 1; i < len(expValues); i++ {
				expVal *= expValues[i]
			}

		case 2:
			expVal = expValues[0]
			for i := 1; i < len(expValues); i++ {
				if expValues[i] < expVal {
					expVal = expValues[i]
				}
			}

		case 3:
			expVal = expValues[0]
			for i := 1; i < len(expValues); i++ {
				if expValues[i] > expVal {
					expVal = expValues[i]
				}
			}

		case 5:
			if expValues[0] > expValues[1] {
				expVal = 1
			}

		case 6:
			if expValues[0] < expValues[1] {
				expVal = 1
			}

		case 7:
			if expValues[0] == expValues[1] {
				expVal = 1
			}
		}
		return
	}
}

func packetDecode(str string, part2 bool) int {
	h := strings.Split(str, "")

	b := bitSet{}
	for _, v := range h {
		val, err := strconv.ParseInt(v, 16, 8)
		if err != nil {
			panic(err)
		}

		b.append4Bits(uint8(val))
	}

	ver, val := b.handlePacket()
	if part2 {
		return int(val)
	}
	return int(ver)
}
