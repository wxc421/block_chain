package blc

import (
	"bytes"
	"crypto/sha256"
	"log/slog"
	"math/big"
)

const targetBit = 16

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	// 256 = 32 * 8
	target = target.Lsh(target, uint(256-targetBit))
	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) Run() ([]byte, int64) {

	var (
		hash    [32]byte
		nonce   = int64(0)
		hashInt = big.NewInt(0)
	)

	for {
		hash = pow.prepareData(nonce)
		hashInt.SetBytes(hash[:])
		if pow.Target.Cmp(hashInt) == 1 {
			// find
			break
		}
		nonce++
	}
	slog.Info("碰撞次数", slog.Attr{
		Key:   "nonce",
		Value: slog.Int64Value(nonce),
	})
	return hash[:], nonce
}

func (pow *ProofOfWork) prepareData(nonce int64) [32]byte {
	tBytes := IntToHex(pow.Block.TimeStamp)
	hBytes := IntToHex(pow.Block.Height)

	bs := bytes.Join([][]byte{
		tBytes,
		hBytes,
		pow.Block.PrevHash,
		pow.Block.Data,
		IntToHex(nonce),
		IntToHex(targetBit),
	}, []byte{})

	hash := sha256.Sum256(bs)
	return hash
}
