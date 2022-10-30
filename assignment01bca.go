package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

type BlockData struct {
	Transactions []string
}
type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
	no          int
}

func CalculateHash(inputBlock *Block) string {
	var temp2 string = ""
	temp := inputBlock.Data.Transactions
	for i := range temp {
		temp2 = temp2 + temp[i]
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(temp2)))
}
func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
	var newBlock = &Block{}
	newBlock.Data = dataToInsert
	newBlock.PrevPointer = nil
	newBlock.PrevHash = ""
	newBlock.CurrentHash = ""
	newBlock.no = 0

	if chainHead != nil {
		newBlock.PrevPointer = chainHead
		newBlock.PrevHash = CalculateHash(chainHead)
		newBlock.CurrentHash = CalculateHash(newBlock)
		newBlock.no = newBlock.PrevPointer.no + 1
	}
	newBlock.CurrentHash = CalculateHash(newBlock)

	return newBlock
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	temp := chainHead
	var i int
	for ; temp != nil; temp = temp.PrevPointer {
		for i = 0; i < len(temp.Data.Transactions); i++ { // v := range temp.Data.Transactions {
			if temp.Data.Transactions[i] == oldTrans {
				temp.Data.Transactions[i] = newTrans
				fmt.Printf("\nBlock Changed!\n")
				return
			}
		}
	}
	fmt.Printf("No old transaction found!")
}
func ListBlocks(chainHead *Block) {
	temp := chainHead
	for ; temp != nil; temp = temp.PrevPointer {
		fmt.Printf("\nCurrent hash: %s\nTransactions: ", temp.CurrentHash)
		for _, v := range temp.Data.Transactions {
			fmt.Printf("%s, ", v)
		}
		fmt.Printf("\n")
	}
}
func VerifyChain(chainHead *Block) {
	temp := chainHead
	for ; temp != nil; temp = temp.PrevPointer {
		if CalculateHash(temp) != temp.CurrentHash {
			fmt.Printf("\nBlockChain Compromised!\n")
			return
		}
	}
	fmt.Printf("\nBlockChain Verified!\n")
}
