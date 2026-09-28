# Working on AdderChess

This is a learning project. Keep the Go engine understandable to its owner. For engine changes, explain the invariant or search idea, why the change is correct, and the measured effect where performance is relevant. Prefer small, reviewable changes over a broad rewrite.

## Where to work

- `src/chess_engine/board` and `gamestate`: representation, FEN, make/undo, and position state.
- `src/chess_engine/move_gen`: legal moves, attacks, pins, and magic tables.
- `src/chess_bot`: evaluation, search, and transposition table.
- `src/server`: HTTP adapter; `interface`: React analysis UI.
- `src/tests`: correctness tests and benchmarks. See `documentation/revival-roadmap.md` for current gaps and priorities.

## Verification

- For Go changes, run relevant tests, then `go test ./...` when available. Add a focused regression for a rule or search bug. For move generation, compare perft counts against known positions before tuning search.
- For UI changes, run `npm ci`, `npm run lint`, and `npm run build` from `interface/` when available. Check the real Go API response and promotion flow, not only the default board.
- For performance work, record the position set, depth or time control, hardware, Go version, baseline and changed results. Report nodes, time, nodes per second, and allocations where applicable. Keep strength results separate from speed results.
- If a toolchain or dependency is unavailable, report which check could not run; do not claim a static inspection is a passing test.

## Design constraints

- Preserve the engine as a standalone Go package. Keep search and evaluation usable without the browser.
- Treat move legality, state restoration, terminal scoring, and draw rules as prerequisites to Elo claims. Do not optimize around a failing correctness test.
- Make the side and unit of every evaluation score explicit (side to move versus White, pawns versus centipawns). Keep UI diagnostics tied to actual engine data.
- Avoid replacing the engine with Stockfish or a chess library; use reference engines as independent comparison tools. Discuss major architecture changes with the owner through concrete evidence and tradeoffs.
