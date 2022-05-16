package utils

import (
	"crypto/sha1"
	// "fmt"
	// "hash/fnv"
)

func GetHash(str string) uint8 {
	// h := fnv.New32a()
	// h.Write([]byte(str))
	// return h.Sum32()

	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	hashVal := uint8(0)

	for _, v := range bs {
		hashVal += uint8(v)
	}
	return hashVal
}
