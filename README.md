# AECH
Blockplain — the 2D (and potentially 3D) evolution of blockchain, forming the core architectural idea behind AECH.


if you run the blockchain, it will print the details of your blocks in a structured way:
Block at [0, 0] - Hash: <genesis_block_hash>
  Previous Block Hash (PrevX, PrevY): , 
  Data: [Genesis]
  Context: Root
  Proof: proof_genesis

Block at [1, 0] - Hash: <proof_1_block_hash>
  Previous Block Hash (PrevX, PrevY): <genesis_block_hash>, 
  Data: [Alice → Bob, Bob → Charlie]
  Context: TimeLayer
  Proof: proof_1
