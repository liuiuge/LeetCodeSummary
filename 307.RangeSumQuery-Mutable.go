package code

// https://leetcode.com/problems/range-sum-query-mutable/

type iNumArray struct {
	arr []int
}

func sConstructor(nums []int) iNumArray {
	var arr1 iNumArray
	arr1.arr = make([]int, 0, len(nums))
	for _, i := range nums {
		arr1.arr = append(arr1.arr, i)
	}
	return arr1
}

func (this *iNumArray) Update(i int, val int) {
	if i < len(this.arr) {
		this.arr[i] = val
	}
	return
}

func (this *iNumArray) SumRange(i int, j int) int {
	sum := 0
	if i > j {
		i, j = j, i
	}
	for k := i; k < len(this.arr) && k <= j; k++ {
		sum += this.arr[k]
	}
	return sum
}
