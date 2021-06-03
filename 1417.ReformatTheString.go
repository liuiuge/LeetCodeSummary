package code

// https://leetcode.com/problems/reformat-the-string/

func reformat(s string) string {
	result := make([]byte, 0, len(s))
	bs := []byte(s)
	dlist, alist := make([]int, 0, len(s)), make([]int, 0, len(s))
	for i := range bs {
		if isdg(bs[i]) {
			dlist = append(dlist, i)
		} else {
			alist = append(alist, i)
		}
	}
	diff := len(dlist) - len(alist)
	if diff*diff <= 1 {
		if len(dlist) > len(alist) {
			dlist, alist = alist, dlist
		}
		for i := range alist {
			result = append(result, bs[dlist[i]])
			result = append(result, bs[alist[i]])
		}
		if len(dlist) < len(alist) {
			result = append(result, bs[dlist[len(alist)-1]])
		}

	}
	return string(result)
}

func isdg(b byte) bool {
	if int(b) >= 48 && int(b) <= 57 {
		return true
	}
	return false
}
