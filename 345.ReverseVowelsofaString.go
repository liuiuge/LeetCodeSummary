package code

// https://leetcode.com/problems/reverse-vowels-of-a-string/

func reverseVowels(s string) string {
	var check_bit struct{}
	vo_map := map[rune]struct{}{
		'a': check_bit,
		'e': check_bit,
		'i': check_bit,
		'o': check_bit,
		'u': check_bit,
		'A': check_bit,
		'E': check_bit,
		'I': check_bit,
		'O': check_bit,
		'U': check_bit,
	}
	locs := make([]int, 0, len(s))
	ns := []rune(s)
	for i := range ns {
		if _, ok := vo_map[ns[i]]; ok {
			locs = append(locs, i)
		}
	}
	for i := 0; i < len(locs)/2; i++ {
		ns[locs[i]], ns[locs[len(locs)-1-i]] = ns[locs[len(locs)-1-i]], ns[locs[i]]
	}
	return string(ns)
}
