package prefix_algoritms

type CharStat map[rune]int

func NewCharStat(text string) CharStat {
	res := CharStat{}

	for _, ch := range text {
		res[ch]++
	}

	return res
}
