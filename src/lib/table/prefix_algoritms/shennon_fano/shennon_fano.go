package shennon_fano

import (
	"archiver/src/lib/table"
	"archiver/src/lib/table/prefix_algoritms"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Generator struct {
}

type encodingTable map[rune]code

type code struct {
	Char     rune
	Quantity int
	Bits     uint32
	Size     int
}

func NewGenerator() Generator {
	return Generator{}
}

func (g Generator) NewTable(text string) table.EncodingTable {
	stat := prefix_algoritms.NewCharStat(text)

	codeTable := build(stat)
	return codeTable.Export()
}

func (et encodingTable) Export() table.EncodingTable {
	res := make(table.EncodingTable)

	for i, v := range et {
		byteString := fmt.Sprintf("%b", v.Bits)

		if lenDiff := v.Size - len(byteString); lenDiff > 0 {
			byteString = strings.Repeat("0", lenDiff) + byteString
		}

		res[i] = byteString
	}

	return res
}

func build(stat prefix_algoritms.CharStat) encodingTable {
	codes := make([]code, 0, len(stat))

	for ch, qty := range stat {
		codes = append(codes, code{
			Char:     ch,
			Quantity: qty,
		})
	}

	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity == codes[j].Quantity {
			return codes[i].Char < codes[j].Char
		}
		return codes[i].Quantity > codes[j].Quantity
	})

	assignCodes(codes)

	res := encodingTable{}

	for _, v := range codes {
		res[v.Char] = v
	}

	return res
}

func assignCodes(codes []code) {
	if len(codes) < 2 {
		return
	}

	divider := bestDividerPosition(codes)

	for i := 0; i < len(codes); i++ {
		codes[i].Bits <<= 1
		codes[i].Size++
		if i >= divider {
			codes[i].Bits |= 1
		}
	}

	assignCodes(codes[:divider])
	assignCodes(codes[divider:])
}

func bestDividerPosition(codes []code) int {
	minDiff := math.MaxInt
	totalSum := 0
	bestPos := 0

	for _, c := range codes {
		totalSum += c.Quantity
	}

	leftSum := 0
	for i, c := range codes {
		leftSum += c.Quantity
		rightSum := totalSum - leftSum

		if int(math.Abs(float64(rightSum-leftSum))) < minDiff {
			minDiff = int(math.Abs(float64(rightSum - leftSum)))
		} else {
			bestPos = i
			break
		}
	}
	return bestPos
}
