package blc

import (
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
	defer blockChain.Close()
	blockChain.AddBlock([]byte("alice send 10 btc to bob"))

	blockChain.AddBlock([]byte("alice send 10 btc to bob"))
	blockChain.PrintBlockChain()

	// for _, block := range blockChain.Blocks {
	// 	fmt.Printf("prevHash:%x,currentHash:%x\n", block.PrevHash, block.Hash)
	// }
}
