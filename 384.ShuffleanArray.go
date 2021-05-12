package code

import (
	"math/rand"
	"time"
)

// https://leetcode.com/problems/shuffle-an-array/

type Solution struct {
	raw []int
}

func SSConstructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{
		raw: nums,
	}

}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.raw

}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	if len(this.raw) < 2 {
		return this.raw
	}
	x := make([]int, len(this.raw))
	for i := range this.raw {
		x[i] = this.raw[i]
	}
	for i := range x {
		tmp := len(x) - rand.Intn(len(x)) - 1
		x[i], x[tmp] = x[tmp], x[i]
	}
	return x
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
