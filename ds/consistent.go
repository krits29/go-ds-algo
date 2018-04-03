package main

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type hashList []uint32

type Ring struct {
	circle        map[uint32]string // all members including replicas
	members       map[string]bool   // unique members
	sortedHashes  hashList          // list of all the members sorted on hashId
	replicasCount int               // count of replicas needed for each member
	count         int               // size of the ring
	sync.RWMutex
}

var error = errors.New("new error")

// sort interface implementation
func (h hashList) Len() int { return len(h) }

func (h hashList) Less(i, j int) bool { return h[i] < h[j] }

func (h hashList) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (r *Ring) Add(node string) {
	r.Lock()
	defer r.Unlock()
	r.add(node)
}

func nodeKey(node string, i int) string {
	return strconv.Itoa(i) + node
}

func hashKey(key string) uint32 {
	h := 0
	for _, c := range key {
		h = h*31 + int(c)
	}
	//return uint32(h)
	return crc32.ChecksumIEEE([]byte(key))
}

func (r *Ring) add(node string) {
	// create required replicas
	for i := 0; i < r.replicasCount; i++ {
		r.circle[hashKey(nodeKey(node, i))] = node
	}
	// add to the members list
	r.members[node] = true
	r.count++
	r.updateSortedHashes()
}

func (r *Ring) Remove(node string) {
	r.Lock()
	defer r.Unlock()

	r.remove(node)
}

func (r *Ring) remove(node string) {
	// create required replicas
	for i := 0; i < r.replicasCount; i++ {
		delete(r.circle, hashKey(nodeKey(node, i)))
	}
	// add to the members list
	delete(r.members, node)
	r.count--
	r.updateSortedHashes()

}

func (r *Ring) updateSortedHashes() {
	// get the latest hashes
	hashes := make(hashList, len(r.sortedHashes))
	for k := range r.circle {
		hashes = append(hashes, k)
	}
	// sort
	sort.Sort(hashes)
	//replace
	r.sortedHashes = hashes
}

func (r *Ring) Get(key string) string {
	r.Lock()
	defer r.Unlock()

	k := hashKey(key)
	i := r.search(k)
	return r.circle[r.sortedHashes[i]]
}

func (r *Ring) search(k uint32) (i int) {
	f := func(x int) bool {
		return r.sortedHashes[x] > k
	}
	i = sort.Search(len(r.sortedHashes), f)
	return
}

func NewRing() *Ring {
	c := new(Ring)
	c.replicasCount = 20
	c.circle = make(map[uint32]string)
	c.members = make(map[string]bool)
	return c
}

func main() {
	c := NewRing()
	c.Add("cacheA")
	c.Add("cacheB")
	c.Add("cacheC")
	users := []string{"user_yogi", "user_yogi1", "user_omar", "user_bunny", "user_stringer"}
	for _, u := range users {
		server := c.Get(u)
		fmt.Printf("%s => %s\n", u, server)
	}
}
