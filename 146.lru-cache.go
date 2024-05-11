/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */
package main

import "fmt"

// @lc code=start
type LRUCache struct {
	cache    *ListNode
	key      int
	capacity int
}
type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{&ListNode{0, 0, nil}, 0, capacity}
}

func (this *LRUCache) Get(key int) int {
	head := this.cache
	this.cache = this.cache.Next
	for head != this.cache {
		if this.cache.Key == key {
			return this.cache.Val
		}
		this.cache = this.cache.Next
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if len(this.cache) == this.capacity {
		delete(this.cache, this.key)
		fmt.Println(key, value, this)
	}
	this.cache[key] = value
	this.key = key
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
