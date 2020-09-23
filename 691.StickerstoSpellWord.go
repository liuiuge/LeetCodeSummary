package code

// https://leetcode.com/problems/stickers-to-spell-word/

type ms_counter struct {
	result      int
	tgt_cnt     [26]int
	sticker_cnt map[int][26]int
}

func minStickers(stickers []string, target string) int {
	tgt_cnt := s_counter(target)
	sticker_cnt := make(map[int][26]int, 0)
	for idx, sticker := range stickers {
		need_cnt := [26]int{}
		temp_cnt := s_counter(sticker)

		for i := range temp_cnt {
			need_cnt[i] = tgt_cnt[i]
			if temp_cnt[i] < need_cnt[i] {
				need_cnt[i] = temp_cnt[i]
			}
		}
		sticker_cnt[idx] = need_cnt

	}
	for i := range sticker_cnt {
		for j := range sticker_cnt {

			if i != j {
				contain := true
				for k := range sticker_cnt[i] {
					if sticker_cnt[i][k] < sticker_cnt[j][k] {
						contain = false
					}
				}
				if contain {
					delete(sticker_cnt, j)
				}
			}
		}
	}
	c := &ms_counter{
		result:      len(target),
		tgt_cnt:     tgt_cnt,
		sticker_cnt: sticker_cnt,
	}

	c.search(0)
	if c.result < len(target) {
		return c.result
	}
	return -1

}

func s_counter(sticker string) [26]int {
	ret := [26]int{}
	for _, v := range sticker {
		ret[v-'a']++
	}
	return ret
}

func (c *ms_counter) search(ans int) {
	if ans >= c.result {
		return
	}
	if len(c.sticker_cnt) < 1 {
		for i := range c.tgt_cnt {
			if c.tgt_cnt[i] > 0 {
				return
			}
		}
		c.result = ans
		return
	}
	temp, pop_key := [26]int{}, 0
	for k, v := range c.sticker_cnt {
		temp = v
		pop_key = k
		break
	}
	delete(c.sticker_cnt, pop_key)
	used := 0

	for i, x := range c.tgt_cnt {
		if temp[i] > 0 {
			t_used := (x-1)/temp[i] + 1
			if t_used > used {
				used = t_used
			}

		}

	}
	if used < 0 {
		used = 0
	}
	for i := range c.tgt_cnt {
		c.tgt_cnt[i] -= used * temp[i]
	}
	c.search(ans + used)
	for i := used - 1; i > -1; i-- {
		for j := range temp {
			c.tgt_cnt[j] += temp[j]
		}
		c.search(ans + i)
	}
	c.sticker_cnt[pop_key] = temp
}
