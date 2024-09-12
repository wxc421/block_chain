package main

import "bkc/blc"

func main() {
	blockChain := blc.NewBlockChain()
	defer blockChain.Close()
	blockChain.AddBlock([]byte("alice send 10 btc to bob"))

	blockChain.AddBlock([]byte("alice send 10 btc to bob"))
	blockChain.PrintBlockChain()

}
