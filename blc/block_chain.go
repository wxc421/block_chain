package blc

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"log/slog"
	"os"
)

const (
	// DB_NAME BlockChain
	DB_NAME = "block.db"
	// BLOCK_TABLE_NAME BlockChain
	BLOCK_TABLE_NAME = "blocks"
	// LATEST_BLOCK_HASH
	LATEST_BLOCK_HASH = "latest_block_hash"
)

// BlockChain 区块链
type BlockChain struct {
	DB              *bolt.DB
	latestBlockHash []byte
}

type BlockChainIterator struct {
	DB              *bolt.DB
	latestBlockHash []byte
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{
		DB:              bc.DB,
		latestBlockHash: bc.latestBlockHash,
	}
}

func (bci *BlockChainIterator) Next() *Block {
	var (
		block *Block
	)
	if bci.latestBlockHash == nil {
		return nil
	}

	err := bci.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCK_TABLE_NAME))
		if b == nil {
			return errors.New("db is nil")
		}
		hash := bci.latestBlockHash
		block = DeSerialize(b.Get(hash))
		bci.latestBlockHash = block.PrevHash
		return nil
	})
	if err != nil {
		return nil
	}
	return block
}

func NewBlockChain() *BlockChain {
	bc := &BlockChain{}
	if err := bc.initBoltDB(); err != nil {
		return nil
	}
	return bc
}

func (bc *BlockChain) Close() {
	if bc.DB != nil {
		bc.DB.Close()
	}
}

func (bc *BlockChain) initBoltDB() error {
	// init db
	db, err := bolt.Open(DB_NAME, os.ModePerm, nil)
	if err != nil {
		slog.Error("init db error", err)
		return err
	}
	bc.DB = db
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BLOCK_TABLE_NAME))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		// if block exists
		if latestBlockHash := b.Get([]byte(LATEST_BLOCK_HASH)); latestBlockHash != nil {
			bc.latestBlockHash = latestBlockHash
			return nil
		}
		// create genesis block
		genesisBlock := NewGenesisBlock([]byte("init block chain"))
		if err = b.Put(genesisBlock.Hash, genesisBlock.Serialize()); err != nil {
			return err
		}
		// save latest block hash
		bc.latestBlockHash = genesisBlock.Hash
		if err = b.Put([]byte(LATEST_BLOCK_HASH), bc.latestBlockHash); err != nil {
			return err
		}
		return err
	})
	if err != nil {
		slog.Error("init db error", err)
		return err
	}
	slog.Info("create db success")
	return nil
}

func (bc *BlockChain) AddBlock(data []byte) error {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCK_TABLE_NAME))
		if b == nil {
			return errors.New("db is nil")
		}
		latestBlock := DeSerialize(b.Get(bc.latestBlockHash))
		block := NewBlock(latestBlock.Hash, latestBlock.Height+1, data)
		if err := b.Put(block.Hash, block.Serialize()); err != nil {
			return err
		}
		// save latest block hash
		bc.latestBlockHash = block.Hash
		if err := b.Put([]byte(LATEST_BLOCK_HASH), bc.latestBlockHash); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (bc *BlockChain) PrintBlockChain() {

	iterator := bc.Iterator()
	for block := iterator.Next(); block != nil; block = iterator.Next() {
		fmt.Printf("--------------------------------------------------------------\n")
		fmt.Printf("\tblock Hash: %x\n", block.Hash)
		fmt.Printf("\tblock PrevHash: %x\n", block.PrevHash)
		fmt.Printf("\tblock Data: %v\n", string(block.Data))
		fmt.Printf("\tblock Height: %v\n", block.Height)
	}
}
