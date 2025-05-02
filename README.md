# AECH
Blockplain — the 2D (and potentially 3D) evolution of blockchain, forming the core architectural idea behind AECH.

# 🧱 AECH: BlockPlain - A 2D Modular Blockchain Protocol

**BlockPlain** is a novel blockchain architecture where blocks are organized in a 2D grid instead of a linear chain. It enables **modular execution**, **native cross-chain communication**, and **parallel processing** of transactions, making it a highly scalable and interoperable blockchain foundation.

---

## 🚀 Features

- 📐 **2D Block Grid**: Blocks are placed in a two-dimensional coordinate plane (x, y) for modularity and spatial separation.
- 🔁 **Cross-Chain Communication**: Supports native interoperability between planes (horizontal and vertical chains).
- 🧠 **Pluggable Consensus**: Modular consensus validation (PoS, PBFT, etc.) with extendable logic.
- 🔍 **Explorer Integration**: Visualize blocks and their placement in the plane with rich metadata.
- 🧰 **Lightweight Golang Core**: Built in Go for performance and concurrency.

---

## 🗂️ Project Structure

```bash
AECH/
│
├── block/         # Block structure & hashing logic
├── bridge/        # Adapter for cross-chain interoperability (WIP)
├── consensus/     # Pluggable consensus logic
├── crosschain/    # Cross-chain transaction handling
├── explorer/      # CLI-based explorer for viewing plane state
├── plane/         # The core 2D block grid logic
├── txpool/        # Transaction pool manager with processing
├── main.go        # Entry point
├── go.mod         # Go module definition
├── test/          # Basic test stub
└── frontend/      # React + TypeScript frontend UI (in progress)
