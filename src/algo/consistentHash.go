package algo

import (
	"fmt"
	"math/rand"
	"sort"
)

type Node struct {
	id        uint32 // hash
	url       string
	strHashed string // url+some random no if collission
}

func GetNode(url string) *Node {
	return &Node{GetHash(url), url, url}
}

type ConsistentHash struct {
	arr  []uint32
	dict map[uint32]Node
}

func GetConsistetnHash() *ConsistentHash {
	return &ConsistentHash{[]uint32{}, map[uint32]Node{}}
}

func (c *ConsistentHash) Add(node *Node) {
	// check if hash for that node already exist
	for {
		if _, exist := c.dict[node.id]; exist {
			node.strHashed += fmt.Sprint(rand.Intn(100))
			node.id = GetHash(node.strHashed)
		} else {
			break
		}
	}

	// add to structure
	c.dict[node.id] = *node
	c.arr = append(c.arr, node.id)
	sort.Slice(c.arr, func(i, j int) bool {
		return c.arr[i] < c.arr[j]
	})
}

// test
func (c *ConsistentHash) GetArray() {
	fmt.Println(c.arr)
}
