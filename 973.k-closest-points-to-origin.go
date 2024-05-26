/*
 * @lc app=leetcode id=973 lang=golang
 *
 * [973] K Closest Points to Origin
 */
package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kClosest([][]int{{1, 3}, {-2, 2}}, 1))
}

// @lc code=start
type Coord struct {
	points []int
	dist   int
}
type PointsHeap []Coord

func kClosest(points [][]int, k int) [][]int {
	h := &PointsHeap{}
	heap.Init(h)
	for _, v := range points {
		heap.Push(h, Coord{
			points: v,
			dist:   v[0]*v[0] + v[1]*v[1],
		})
		fmt.Println(h)
	}
	res := [][]int{}
	for i := 0; i < k; i++ {
		coord := heap.Pop(h).(Coord)
		res = append(res, coord.points)
	}
	return res
}

func (h *PointsHeap) Push(x any) {
	if 

	*h = append(*h, x.(Coord))
}
func (h PointsHeap) Len() int           { return len(h) }
func (h PointsHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h PointsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointsHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
