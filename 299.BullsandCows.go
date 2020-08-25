package code

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	secret_lista := make([]int, len(secret))
	secret_listb := make([]int, len(secret))
	cnta, cntb := 0, 0
	for i := range secret {
		if secret[i] == guess[i] {
			cnta++
		} else {
			secret_lista[int(secret[i])-46]++
			secret_listb[int(guess[i])-46]++
		}
	}
	for i := range secret_lista {
		tmp := secret_lista[i]
		if secret_listb[i] < secret_lista[i] {
			tmp = secret_listb[i]
		}
		cntb += tmp
	}
	return fmt.Sprintf("%dA%dB", cnta, cntb)
}
