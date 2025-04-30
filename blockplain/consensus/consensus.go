package consensus

import (
	"strings"
)

// ValidateProof simulates a simple proof validation.
// You can replace this later with PoS, PBFT, or other mechanisms.
func ValidateProof(proof string) bool {
	// Example: proof must start with "proof_" and not be empty
	if proof == "" {
		return false
	}
	return strings.HasPrefix(proof, "proof_")
}
