package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Block struct {
	X, Y        int
	PrevHashX   string
	PrevHashY   string
	Data        []string
	Context     string
	Proof       string
	Hash        string
}

// Compute a unique SHA-256 hash of the block
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

// Constructor for a new block
func NewBlock(x, y int, txs []string, context, proof, prevHashX, prevHashY string) *Block {
	b := &Block{
		X:         x,
		Y:         y,
		Data:      txs,
		Context:   context,
		Proof:     proof,
		PrevHashX: prevHashX,
		PrevHashY: prevHashY,
	}
	b.Hash = b.ComputeHash()
	return b
}
