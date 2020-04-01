package exthash

import "hash/crc64"

func Crc64(str string) uint64{
	//先建立一個table
	table := crc64.MakeTable(crc64.ECMA)
	//傳入位元組切片和table，返回一個uint64
	return crc64.Checksum([]byte(str), table)
}
