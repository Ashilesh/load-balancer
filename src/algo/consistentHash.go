package algo

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/Ashilesh/load-balancer/logs"
	"github.com/Ashilesh/load-balancer/utils"
)

type serverNode struct {
	id        uint8 // hash
	url       string
	strHashed string // url+some random no if collission
}

func getNode(url string) *serverNode {
	return &serverNode{utils.GetHash(url), url, url}
}

type ConsistentHash struct {
	arr  []uint8
	dict map[uint8]serverNode
}

var consistentHashAlgo *ConsistentHash

func GetConsistetnHash() *ConsistentHash {
	if consistentHashAlgo != nil {
		return consistentHashAlgo
	}
	consistentHashAlgo = &ConsistentHash{[]uint8{}, map[uint8]serverNode{}}
	return consistentHashAlgo
}

func (c *ConsistentHash) Add(url string) {
	node := getNode(url)

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
	if ind, exist := utils.BinarySearch(c.arr, utils.GetHash(modifiedUrl)); exist {
		c.arr = append(c.arr[:ind], c.arr[ind+1:]...)
		delete(c.dict, utils.GetHash(modifiedUrl))
	}
}

func (c *ConsistentHash) GetUrl(clientIP string) string {
	ind, _ := utils.BinarySearch(c.arr, utils.GetHash(clientIP))

	if ind < 0 {
		// TODO: change this implementation when we listen for server status
		logs.Fatal("0 nodes available")
	}

	node := c.dict[c.arr[ind]]

	return node.url
}
