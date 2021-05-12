package code

// https://leetcode.com/problems/distribute-candies-to-people/

func distributeCandies(candies int, num_people int) []int {
	start, count, cnt, sum, temp := 1, 0, []int{}, 0, 0
	for {
		temp = (start + start + num_people - 1) * num_people / 2
		// fmt.Println(temp, count)
		count += temp
		if candies > count {
			cnt = append(cnt, start)
			sum += start
		} else {
			break
		}
		start += num_people
	}
	left := candies - count + temp
	result := make([]int, num_people)
	// fmt.Println(start, count, cnt, sum, temp, left)
	for i := range result {
		if len(cnt) > 0 {
			result[i] = (cnt[len(cnt)-1] + i + i + 1) * len(cnt) / 2
		} else {
			result[i] = 0
		}
		if start < left {
			result[i] += start
			left -= start
			start++
		} else {
			result[i] += left
			left = 0
		}
		// fmt.Println(start, left)
	}
	return result
}
