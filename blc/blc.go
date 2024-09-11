package blc

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log/slog"
	"time"
)

type Block struct {
	TimeStamp int64  // 时间戳
	Hash      []byte // 当前区块哈希
	PrevHash  []byte // 前区块哈希
	Height    int64  // 区块高度
	Data      []byte // 交易数据
	Nonce     int64  // 随机值
	Diff      int64  // 难度系数
}

func NewBlock(prevHash []byte, height int64, data []byte) *Block {
	block := &Block{TimeStamp: time.Now().Unix(), PrevHash: prevHash, Height: height, Data: data}

	pow := NewProofOfWork(block)
	// 执行工作量证明算法
	block.Hash, block.Nonce = pow.Run()
	return block
}

// NewGenesisBlock 创世区块
func NewGenesisBlock(data []byte) *Block {
	return NewBlock(nil, 1, data)
}

func (b *Block) SetHash() {

	tBytes := IntToHex(b.TimeStamp)
	hBytes := IntToHex(b.Height)

	bs := bytes.Join([][]byte{
		tBytes,
		hBytes,
		b.PrevHash,
		b.Data,
	}, []byte{})

	hash := sha256.Sum256(bs)
	b.Hash = hash[:]
}

func IntToHex(data int64) []byte {
	buffer := &bytes.Buffer{}
	err := binary.Write(buffer, binary.BigEndian, data)
	if err != nil {
		slog.Warn("IntToHex error:%v", err)
	}
	return buffer.Bytes()
}
