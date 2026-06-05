package huffman

import (
	"archiver/src/lib/table"
	"archiver/src/lib/table/prefix_algoritms"
	"sort"
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

type binTree struct {
	one      *binTree
	zero     *binTree
	priority int
	val      code
}

func (g Generator) NewTable(text string) table.EncodingTable {
	stat := prefix_algoritms.NewCharStat(text)

	codeTable := build(stat)
	_ = codeTable
	//return codeTable.Export()
	return nil
}

//func (et encodingTable) Export() table.EncodingTable {
//	res := make(table.EncodingTable)
//
//	return res
//}

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
			return codes[i].Char > codes[j].Char
		}
		return codes[i].Quantity < codes[j].Quantity
	})

	huffmanBinTree := getHuffmanBinTree(codes)
	res := encodingTable{}
	assignCodes(&res, *huffmanBinTree.zero, 0, 1)
	assignCodes(&res, *huffmanBinTree.one, 1, 1)

	return res
}

func assignCodes(resTable *encodingTable, binTree binTree, code uint32, size int) {
	if binTree.val.Char != 0 {
		binTree.val.Size = size
		binTree.val.Bits = code
		(*resTable)[binTree.val.Char] = binTree.val
	} else {
		assignCodes(resTable, *binTree.one, code<<1+1, size+1)
		assignCodes(resTable, *binTree.zero, code<<1+0, size+1)
	}
}

func getHuffmanBinTree(codes []code) binTree {
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
			//left := binTrees[:i]
			//right := binTrees[i:]
			//left = append(left, newBinTree)
			//left = append(left, right...)
			//binTrees = left
			return append(binTrees[:i], append([]binTree{newBinTree}, binTrees[i:]...)...)
		}
	}

	binTrees = append(binTrees, newBinTree)

	return binTrees
}
