package code

import (
	"strings"
)

// https://leetcode.com/problems/rearrange-spaces-between-words/

func reorderSpaces(text string) string {
	cnt, idx, word := 0, 0, []string{}
	r_text := []rune(text)
	for idx < len(text) {
		if r_text[idx] == ' ' && idx == 0 {
			for r_text[idx] == ' ' {
				cnt++
				idx++
			}
		} else if r_text[idx] != ' ' {
			temp := []rune{}
			for r_text[idx] != ' ' {
				temp = append(temp, r_text[idx])
				idx++
			}
			word = append(word, string(temp))
		} else {
			cnt++
			idx++
		}
	}
	gap, tail := cnt/len(word), cnt%len(word)
	gap_word := ""
	for i := 0; i < gap; i++ {
		gap_word += " "
	}
	tail_word := ""
	for i := 0; i < tail; i++ {
		tail_word += " "
	}
	return strings.Join(word, gap_word) + tail_word

}
