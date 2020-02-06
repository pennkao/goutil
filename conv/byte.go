package conv

import "encoding/binary"

//int64 转 byte big
func Int64ToBytesBig(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

//int64 转 byte little
func Int64ToBytesLittle(i int64) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}
