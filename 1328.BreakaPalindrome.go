package code

// https://leetcode.com/problems/break-a-palindrome/

func breakPalindrome(palindrome string) string {
	if len(palindrome) < 2 {
		return ""
	} else {
		x := []rune(palindrome)
		for i := 0; i < len(palindrome)/2; i++ {
			if x[i] != 'a' {
				x[i] = 'a'
				return string(x)
			}
		}
		x[len(palindrome)-1] = 'b'
		palindrome = string(x)
	}
	return palindrome

}
