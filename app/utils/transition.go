package utils

import "encoding/binary"

func Int64ToByteArray(num int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(buf, num)
	return buf
}

func ByteArrayToInt64(buf []byte) int64 {
	num, _ := binary.Varint(buf)
	return num
}
