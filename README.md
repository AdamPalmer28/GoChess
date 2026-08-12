# AdderChess

AdderChess is a from-scratch chess engine and analysis application. The core engine is written in Go, with a React/Vite browser interface connected through a small JSON HTTP API.

> **Project status:** inactive / archived. This repository is preserved as a portfolio case study rather than a production-ready chess service. The implementation and setup instructions reflect the state of the project when it was actively developed.

## Portfolio snapshot

This project demonstrates the design and integration of:

- a compact, stateful chess engine built around 64-bit bitboards;
- legal move generation, including king-safety constraints and special moves;
- a search bot using negamax-style alpha-beta pruning, quiescence search, and a transposition table;
- a material and positional evaluation pipeline with per-piece analysis data;
- a Go HTTP server that exposes game and analysis state as JSON;
- a React interface for playing through positions and inspecting engine output;
- unit tests and repeatable Go benchmarks for correctness and performance work.

## Architecture

```text
React + Vite UI
        │ JSON over HTTP
        ▼
Go server / GameHost
        │
        ├── GameState and reversible move history
        ├── Bitboard board representation
        ├── Legal move generation and king-safety analysis
        └── Search bot and evaluation pipeline
```

The code is intentionally split into layers so the engine can be exercised independently from the browser application:

```text
src/chess_engine/   board, FEN, game state, move generation, magic tables
src/chess_bot/      search, alpha-beta, transposition table, evaluation
src/server/         HTTP endpoints and JSON data shaping
src/tests/           move-generation and game-state tests/benchmarks
interface/          React/Vite chess board and analysis UI
perf/               benchmark scripts and historical results
```

## Engine capabilities

### Board and game state

- Separate 64-bit bitboards for every piece type and side occupancy.
- FEN position loading for focused engine tests and analysis.
- Compact 16-bit move encoding containing the origin square, destination square, and move type.
- Make/undo support with history for captured pieces, castling rights, en passant state, and previous moves.
- Zobrist hashing for position identity and search-table lookups.
- Detection of check, checkmate, stalemate, and game-over positions.

### Move generation

- Pawn, knight, bishop, rook, queen, and king move generation.
- Castling, en passant, captures, and promotion moves.
- Pin detection and removal of moves that expose the king.
- Dedicated in-check move generation for single and double-check positions.
- Precomputed attack rays for knights, kings, pawns, and sliding pieces.
- Magic-bitboard lookup tables for bishop and rook attacks.

### Search and evaluation

- Alpha-beta pruning using a negamax search structure.
- Quiescence search to extend tactical positions beyond the nominal search depth.
- Transposition table keyed by the Zobrist position hash.
- Search statistics including node count, pruning, transposition-table hits, and successful table lookups.
- Evaluation components for:
  - material balance;
  - pawn structure, including doubled, isolated, backward, and connected pawns;
  - pawn centre control and promotion potential;
  - knight placement and edge penalties;
  - bishop, rook, and queen activity using move rays and x-ray relationships;
  - king safety and check status.

The bot also exposes a structured evaluation breakdown, which allows the UI to show how the overall score was assembled instead of presenting a single opaque number.

## Browser interface

The React/Vite interface provides a visual way to exercise the engine and API:

- interactive board rendering with selectable legal moves;
- last-move and available-move highlighting;
- board flip, new-game, and undo controls;
- move history and engine-ordered move lists;
- evaluation bar and per-piece evaluation breakdown;
- game-state and analysis tabs backed by live server responses.

The interface communicates with the Go server using these endpoints:

| Endpoint | Purpose |
| --- | --- |
| `GET /chessgame` | Current board, legal moves, history, and state flags |
| `POST /move` | Validate and apply a client move |
| `POST /undo` | Revert the most recent move |
| `POST /newgame` | Start a fresh game state |
| `GET /analysis` | Current evaluation breakdown and selected bitboard data |
| `GET /EvalScoreData` | Evaluation data endpoint |
| `GET /BitBoardData` | Bitboard data endpoint |

## Running locally

### Requirements

- Go 1.20 or later
- Node.js and npm

### Start the Go server

From the repository root:

```bash
go run .
```

The server listens on `http://localhost:8080`.

### Start the React interface

In a second terminal:

```bash
cd interface
npm install
npm run dev
```

Open `http://localhost:4000`. The interface expects the Go server to be running on port `8080`.

## Testing and benchmarking

Run the Go test suite:

```bash
go test ./...
```

The tests cover board primitives, piece-specific move generation, pinned pieces, check and checkmate handling, FEN parsing, make/undo behaviour, and game-state transitions.

Run the Go benchmarks:

```bash
go test ./src/tests/... -bench=. -run=^# -benchmem
```

The benchmark package measures game-state preparation, move generation, evaluation components, and best-move search across representative FEN positions. Historical benchmark output is kept in [`perf/benchmark_data`](perf/benchmark_data).

For the frontend:

```bash
cd interface
npm run lint
npm run build
```

## Technical takeaways

The most valuable part of this project is the end-to-end reasoning it required: representing a complex ruleset compactly, maintaining consistent state through make/undo operations, validating legal moves under pins and checks, and carrying engine internals across a clean API boundary into a usable interface. It is a useful snapshot of work spanning algorithms, systems programming, testing, performance investigation, and frontend integration.

## Repository guide

- [`src/chess_engine`](src/chess_engine) — board representation, game state, FEN parsing, move generation, and magic tables
- [`src/chess_bot`](src/chess_bot) — search and evaluation
- [`src/server`](src/server) — Go HTTP API
- [`src/tests`](src/tests) — correctness tests and benchmarks
- [`interface`](interface) — React/Vite UI
- [`documentation`](documentation) — engine and server notes
- [`perf`](perf) — benchmark scripts and recorded results
