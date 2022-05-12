package algo

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/Ashilesh/balancer/src/utils"
)

type Algo interface {
	Add(string)
	Delete(string)
	GetUrl(string) string
}

// TODO: pass algo type in params
func GetAlgo() Algo {
	return GetConsistetnHash()
}

type Node struct {
	id        uint32 // hash
	url       string
	strHashed string // url+some random no if collission
}

func GetNode(url string) *Node {
	return &Node{utils.GetHash(url), url, url}
}

type ConsistentHash struct {
	arr  []uint32
	dict map[uint32]Node
}

func GetConsistetnHash() *ConsistentHash {
	return &ConsistentHash{[]uint32{}, map[uint32]Node{}}
}

func (c *ConsistentHash) Add(url string) {
	node := GetNode(url)

	// check if hash for that node already exist
	for {
		if _, exist := c.dict[node.id]; exist {
			node.strHashed += fmt.Sprint(rand.Intn(100))
			node.id = utils.GetHash(node.strHashed)
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

func (c *ConsistentHash) Delete(modifiedUrl string) {
	if ind, exist := utils.Search(c.arr, utils.GetHash(modifiedUrl)); exist {
		c.arr = append(c.arr[:ind], c.arr[ind+1:]...)
		delete(c.dict, utils.GetHash(modifiedUrl))
	}
}

func (c *ConsistentHash) GetUrl(clientIP string) string {
	ind, _ := utils.Search(c.arr, utils.GetHash(clientIP))

	if ind < 0 {
		panic("0 Nodes available")
	}

	node := c.dict[c.arr[ind]]

	return node.url
}
