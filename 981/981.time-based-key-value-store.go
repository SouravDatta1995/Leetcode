/*
 * @lc app=leetcode id=981 lang=golang
 *
 * [981] Time Based Key-Value Store
 */
package main

import (
	"fmt"
)

func main() {
	var timeMap = Constructor()
	timeMap.Set("foo", "bar", 1)
	timeMap.Set("foo", "bar2", 4)
	fmt.Println(timeMap.Get("foo", 7))
}

// @lc code=start
type TimeMap struct {
	key map[string][]Pair
}

type Pair struct {
	value string
	time  int
}

func Constructor() TimeMap {
	return TimeMap{key: make(map[string][]Pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.key[key]; !ok {
		this.key[key] = []Pair{}
	}
	this.key[key] = append(this.key[key], Pair{time: timestamp, value: value})
	// slices.SortFunc(this.key[key], func(a, b Pair) int {
	// 	return cmp.Compare(a.time, b.time)
	// })

}

func (this *TimeMap) Get(key string, timestamp int) string {
	res := ""
	val, ok := this.key[key]
	if ok {
		l, r := 0, len(val)-1
		for l <= r {
			mid := (l + r) / 2
			if val[mid].time == timestamp {
				return val[mid].value
			} else if val[mid].time < timestamp {
				res = val[mid].value
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
// @lc code=end
