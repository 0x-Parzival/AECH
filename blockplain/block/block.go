package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	X, Y        int
	PrevHashX   string
	PrevHashY   string
	Data        []string
	Context     string
	Proof       string
	Hash        string
	Timestamp   time.Time // New field for timestamp
}

func (b *Block) ComputeHash() string {
	h := sha256.New()
	h.Write([]byte(
		strconv.Itoa(b.X) +
			strconv.Itoa(b.Y) +
			b.PrevHashX +
			b.PrevHashY +
			b.Context +
			b.Proof,
	))
	for _, d := range b.Data {
		h.Write([]byte(d))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func NewBlock(x, y int, txs []string, context string, proof string, prevHashX string, prevHashY string) *Block {
	block := &Block{
		X: x, Y: y,
		Data:      txs,
		Context:   context,
		Proof:     proof,
		PrevHashX: prevHashX,
		PrevHashY: prevHashY,
		Timestamp: time.Now(),  // Add timestamp when the block is created
	}
	block.Hash = block.ComputeHash()
	return block
}
