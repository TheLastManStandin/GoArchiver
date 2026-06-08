package huffman

import (
	"archiver/src/lib/table"
	"archiver/src/lib/table/prefix_algoritms"
	"slices"
	"sort"
)

type Generator struct {
}

func NewGenerator() Generator {
	return Generator{}
}

type binTree struct {
	one      *binTree
	zero     *binTree
	priority int
	val      prefix_algoritms.Code
}

func (g Generator) NewTable(text string) table.EncodingTable {
	stat := prefix_algoritms.NewCharStat(text)

	codeTable := build(stat)
	return codeTable.Export()
}

func build(stat prefix_algoritms.CharStat) prefix_algoritms.EncodingTable {
	codes := make([]prefix_algoritms.Code, 0, len(stat))

	for ch, qty := range stat {
		codes = append(codes, prefix_algoritms.Code{
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

	huffmanBinTree := getHuffmanBinTree(codes)
	res := prefix_algoritms.EncodingTable{}
	assignCodes(&res, *huffmanBinTree.zero, 0, 1)
	assignCodes(&res, *huffmanBinTree.one, 1, 1)

	return res
}

func assignCodes(resTable *prefix_algoritms.EncodingTable, binTree binTree, code uint32, size int) {
	if binTree.val.Char != 0 {
		binTree.val.Size = size
		binTree.val.Bits = code
		(*resTable)[binTree.val.Char] = binTree.val
	} else {
		assignCodes(resTable, *binTree.one, code<<1+1, size+1)
		assignCodes(resTable, *binTree.zero, code<<1+0, size+1)
	}
}

func getHuffmanBinTree(codes []prefix_algoritms.Code) binTree {
	//sort.Slice(codes, func(i, j int) bool {
	//	if codes[i].Quantity == codes[j].Quantity {
	//		return codes[i].Char < codes[j].Char
	//	}
	//	return codes[i].Quantity < codes[j].Quantity
	//})
	binTrees := make([]binTree, len(codes))

	for i, v := range codes {
		binTrees[i] = binTree{
			val:      v,
			priority: v.Quantity,
		}
	}

	for len(binTrees) > 1 {
		newBinTree := binTree{
			one:      &binTrees[1],
			zero:     &binTrees[0],
			priority: binTrees[0].priority + binTrees[1].priority,
		}

		binTrees = insertNewBinTree(binTrees[2:], newBinTree)
	}

	finalBinTree := binTrees[0]

	return finalBinTree
}

func insertNewBinTree(binTrees []binTree, newBinTree binTree) []binTree {
	for i, val := range binTrees {
		if val.priority >= newBinTree.priority {
			return slices.Insert(binTrees, i, newBinTree)
		}
	}

	binTrees = append(binTrees, newBinTree)

	return binTrees
}
