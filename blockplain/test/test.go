package test

import (
	"testing"
	"blockplain/block"
	"blockplain/consensus"
	"blockplain/plane"
)

// TestBlockHash checks if block hash is computed consistently
func TestBlockHash(t *testing.T) {
	b := block.NewBlock(0, 0, []string{"A → B"}, "TestCtx", "test_proof", "", "")
	expected := b.ComputeHash()
	if b.Hash != expected {
		t.Errorf("Expected hash %s, got %s", expected, b.Hash)
	}
}

// TestConsensusProof checks proof validation logic
func TestConsensusProof(t *testing.T) {
	valid := consensus.ValidateProof("not_empty")
	if !valid {
		t.Error("Expected proof to be valid")
	}

	invalid := consensus.ValidateProof("")
	if invalid {
		t.Error("Expected empty proof to be invalid")
	}
}

// TestPlaneInsertion verifies block addition and retrieval
func TestPlaneInsertion(t *testing.T) {
	p := plane.NewPlane()
	b := block.NewBlock(1, 2, []string{"Tx1"}, "Layer", "proof", "", "")
	p.AddBlock(1, 2, b)

	ret := p.GetBlock(1, 2)
	if ret == nil || ret.Hash != b.Hash {
		t.Error("Failed to retrieve correct block from plane")
	}
}
