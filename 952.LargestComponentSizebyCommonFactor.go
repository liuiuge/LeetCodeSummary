package code

// https://leetcode.com/problems/largest-component-size-by-common-factor/

type Dsu struct {
	MxSize int
	Parent []int
	Size   []int
}

func NewDsu(size int) *Dsu {
	x := &Dsu{
		MxSize: 1,
		Parent: make([]int, 100005),
		Size:   make([]int, 100005),
	}
	for i := range x.Parent {
		x.Parent[i] = i
		x.Size[i] = 1
	}
	return x
}

func (d *Dsu) findInLocal(x int) int {
	if x != d.Parent[x] {
		d.Parent[x] = d.findInLocal(d.Parent[x])
	}
	return d.Parent[x]
}

func (d *Dsu) merge(x, y int) {
	x = d.findInLocal(x)
	y = d.findInLocal(y)
	if x == y {
		return
	}
	if d.Size[x] < d.Size[y] {
		x, y = y, x
	}
	d.Parent[y] = x
	d.Size[x] += d.Size[y]
	if d.MxSize < d.Size[x] {
		d.MxSize = d.Size[x]
	}
}

func (d *Dsu) ans() int {
	return d.MxSize
}

func largestComponentSize(nums []int) int {

	high := nums[0]
	dict_n := make(map[int]struct{}, len(nums))
	for i := range nums {
		dict_n[nums[i]] = struct{}{}
		if nums[i] > high {
			high = nums[i]
		}
	}
	d := NewDsu(high + 1)
	visited := make(map[int]struct{})
	for i := 2; 2*i <= high; i++ {
		if _, vok := visited[i]; !vok {
			last := -1
			for j := i * 2; j <= high; j += i {
				visited[j] = struct{}{}
				if last == -1 {
					if _, jok := dict_n[j]; jok {
						last = j
					}
				} else if _, jok := dict_n[j]; jok {
					d.merge(last, j)
				}
			}
		}
	}
	return d.ans()
}

func LargestComponentSize(nums []int) int {
	return largestComponentSize(nums)
}
