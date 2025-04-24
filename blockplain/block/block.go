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
