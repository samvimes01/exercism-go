package yacht

import "sort"

func Score(dice []int, category string) int {
	var n int
	switch category {
	case "ones":
		n = 1
	case "twos":
		n = 2
	case "threes":
		n = 3
	case "fours":
		n = 4
	case "fives":
		n = 5
	case "sixes":
		n = 6
	case "full house":
		return FullHouse(dice)
	case "four of a kind":
		return FourOfAKind(dice)
	case "little straight":
		return Straight(dice, 1)
	case "big straight":
		return Straight(dice, 2)
	case "choice":
		return Sum(dice)
	case "yacht":
		return Yacht(dice)
	}
	return ScoreMult(dice, n)
}

func FourOfAKind(dice []int) int {
	sort.Ints(dice)
	if dice[0] == dice[3] {
		return Sum(dice[0:4])
	}
	if dice[1] == dice[4] {
		return Sum(dice[1:5])
	}
	return 0
}

func FullHouse(dice []int) int {
	dict := make(map[int]int)
	for _, d := range dice {
		dict[d]++
	}
	if len(dict) != 2 {
		return 0
	}
	cnt := make([]int, 0)
	for _, v := range dict {
		cnt = append(cnt, v)
	}
	if (cnt[0] == 3 && cnt[1] == 2) || (cnt[0] == 2 && cnt[1] == 3) {
		return Sum(dice)
	}
	return 0
}

func Straight(dice []int, startFrom int) int {
	straights := make([]int, 5)
	for i := range straights {
		straights[i] = startFrom + i
	}
	sort.Ints(dice)
	for i, d := range dice {
		if d != straights[i] {
			return 0
		}
	}
	return 30
}

func Yacht(dice []int) int {
	for i := 1; i < len(dice); i++ {
		if dice[i] != dice[i-1] {
			return 0
		}
	}
	return 50
}

func Sum(dice []int) int {
	sum := 0
	for _, d := range dice {
		sum += d
	}
	return sum
}

func ScoreMult(dice []int, num int) int {
	sum := 0
	for _, d := range dice {
		if d == num {
			sum += num
		}
	}
	return sum
}
