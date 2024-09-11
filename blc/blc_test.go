package blc

import (
	"fmt"
	"testing"
)

func TestNewBlock(t *testing.T) {
	type args struct {
		prevHash []byte
		height   int64
		data     []byte
	}
	tests := []struct {
		name string
		args args
		want *Block
	}{
		{
			name: "NewBlock",
			args: args{
				prevHash: nil,
				height:   1,
				data:     []byte("the first block testing"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := NewBlock(tt.args.prevHash, tt.args.height, tt.args.data)
			t.Log(block)
			// if got := block; !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NewBlock() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestBlock_Hash(t *testing.T) {
	blockChain := NewBlockChain()
	lastBlock := blockChain.LastBlock()
	blockChain.AddBlock(NewBlock(lastBlock.Hash, lastBlock.Height+1, []byte("alice send 10 btc to bob")))

	lastBlock = blockChain.LastBlock()
	blockChain.AddBlock(NewBlock(lastBlock.Hash, lastBlock.Height+1, []byte("bob send 5 btc to alice")))

	for _, block := range blockChain.Blocks {
		fmt.Printf("prevHash:%x,currentHash:%x\n", block.PrevHash, block.Hash)
	}
}
