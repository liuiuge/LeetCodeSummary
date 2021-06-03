package code

import (
	"fmt"
)

type MapSum struct {
	Kvmap map[string]int
	Pfmap map[string]int
}

/** Initialize your data structure here. */
func MapSumConstructor() MapSum {
	return MapSum{
		Kvmap: make(map[string]int, 0),
		Pfmap: make(map[string]int, 0),
	}
}

func (this *MapSum) updatePf(key string, val int) {
	for i := 1; i < len(key); i++ {
		ret, ok := this.Pfmap[key[:i]]
		if ok {
			this.Pfmap[key[:i]] = ret + val
		} else {
			this.Pfmap[key[:i]] = val
		}
	}
	fmt.Println(this.Pfmap)
}

func (this *MapSum) Insert(key string, val int) {
	u_val := val
	ret, ok := this.Kvmap[key]
	if ok {
		u_val = val - ret
	}
	this.Kvmap[key] = val
	this.updatePf(key, u_val)
}

func (this *MapSum) Sum(prefix string) int {
	r_ret, ok := this.Pfmap[prefix]
	if !ok {
		return 0
	}
	return r_ret
}
