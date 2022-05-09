package algo

import (
	"hash/fnv"
)

func GetHash(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}
