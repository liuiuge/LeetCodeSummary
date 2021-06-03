package code

// https://leetcode.com/problems/valid-parenthesis-string/

func checkValidString(s string) bool {
	bs := []byte(s)
	cl, cr := 0, 0
	for i := range bs {
		if bs[i] == '(' {
			cl++
			cr++
		} else if bs[i] == ')' {
			cl--
			cr--
		} else {
			cl++
			cr--
		}
		if cl < 0 {
			return false
		}
		if cr < 0 {
			cr = 0
		}
	}
	return true
}
