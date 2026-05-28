package shennon_fano

import (
	"archiver/src/lib/compression/vlc/table"
	"math"
	"sort"
)

type Generator struct {
}

type charStat map[rune]int

type encodingTable map[rune]code

type code struct {
	Char     rune
	Quantity int
	Bit      uint32
	Size     int
}

func NewGenerator() Generator {
	return Generator{}
}

func (g Generator) NewTable(text string) table.EncodingTable {
	//stat := newCharStat(text)

	return nil
}

func build(stat charStat) encodingTable {
	codes := make([]code, 0, len(stat))

	for ch, qty := range stat {
		codes = append(codes, code{
			Char:     ch,
			Quantity: qty,
		})
	}

	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity == codes[j].Quantity {
			return codes[i].Char > codes[j].Char
		}
		return codes[i].Quantity < codes[j].Quantity
	})

	assignCodes(codes)

	return nil
}

func assignCodes(codes []code) {
	if len(codes) < 2 {
		return
	}

}

func bestDividerPosition(codes []code) int {
	minDiff := math.MaxInt
	totalSum := 0

	for _, c := range codes {
		totalSum += c.Quantity
	}

	leftSum := codes[0].Quantity
	for i, c := range codes {
		rightSum := totalSum - leftSum

		if int(math.Abs(float64(rightSum-leftSum))) < minDiff {
			minDiff = int(math.Abs(float64(rightSum - leftSum)))
		} else {
			return i - 1
		}
		leftSum += c.Quantity
	}
	return len(codes) - 1
}

func newCharStat(text string) charStat {
	res := charStat{}

	for _, ch := range text {
		res[ch]++
	}

	return res
}
