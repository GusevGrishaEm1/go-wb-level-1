package main

func setBit(num int64, pos int) int64 {
	mask := int64(1) << pos
	num |= mask
	return num
}

func unsetBit(num int64, pos int) int64 {
	mask := ^(int64(1) << pos)
	num &= mask
	return num
}