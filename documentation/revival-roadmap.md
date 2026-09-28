# AdderChess revival audit — 28 September 2026

This is a source review, not a verified strength or correctness rating. The Go toolchain was unavailable in the review environment, and the frontend dependencies were not installed, so `go test ./...`, `npm run lint`, and `npm run build` could not complete. Recheck these findings when the project is runnable.

## Standing

| Area | Current state | Main gap |
| --- | --- | --- |
| Documentation | The root README maps the architecture and setup. Feature and todo notes describe intent. | No explanation of state invariants, search score conventions, API contract, or known bugs. Some notes describe planned features as if present. |
| Rules engine | Bitboards, FEN loading, move generation, make/undo, castling, en passant, promotion, checkmate and stalemate paths exist. | No perft reference suite; draw rules are absent; state restoration and special moves need adversarial checks. |
| Bot | Evaluation, fixed depth negamax/alpha-beta, quiescence and an in-memory transposition table exist. The CLI invokes the bot. | Search has apparent correctness errors. No iterative deepening, time management, UCI adapter, or measured Elo. |
| Tests | 23 Go test functions and 3 benchmark functions cover board operations, move types and selected FENs. | No bot or HTTP tests, no frontend tests, no perft suite or automated CI gate. Existing benchmarks do not establish playing strength. |
| UI | React/Vite has a board, move list, evaluation breakdown, highlights and basic controls. | Bot/settings/game panels are placeholders; no bot endpoint or promotion choice. Some controls and displayed data are unreliable. |

### Findings to verify first

- `src/chess_bot/alpha_beta.go`: terminal nodes update `alpha` but continue, mate score changes sign with board color, and quiescence uses stand-pat even in check while searching only moves selected by a score threshold. These can change move choice, not just speed.
- `src/chess_bot/transposition.go` and `alpha_beta.go`: entries have no bound type; values from cutoffs are reused as exact scores. Depth and score perspective need explicit tests.
- `src/chess_engine/gamestate/make_move.go` and `undo.go`: the halfmove clock is decremented on undo but never updated on make; undo restores only the mover's castling rights even when an opponent rook was captured. Draw adjudication is limited to no legal moves (`gamestate.go`).
- `src/server/gameData.go`: the new-game handler replaces a captured local host pointer, while the analysis handlers retain the original host. An invalid move writes an error and then tries to write game JSON. The move request does not preserve the promotion choice.
- `interface/src/Components/chess/board_ui.jsx`, `board.jsx`, and `.api/api.jsx`: promotion is not selectable; board flip reverses ranks but not files and mutates a shared row array; move/undo/new-game wrappers do not return their fetch promise, so analysis refresh can race with the state change. The board evaluation bar uses a fixed score.

## Work items

Each checkbox is intended to be a separate, reviewable agent task with a runnable acceptance check.

### Restore the baseline

- [ ] **R1 — Reproduce a clean run.** Install a supported Go toolchain and frontend dependencies, run Go tests and UI lint/build, start both services, and document exact versions, commands, failures and one browser smoke test. Fix only setup blockers in this item.
- [ ] **R2 — Establish a correctness oracle.** Add perft for the initial position and several castling, en passant, promotion, pin and check positions, with trusted reference counts and a divide mode for debugging. Require it to pass before search tuning.
- [ ] **R3 — Make state round trips exact.** Test make/undo across every special move and rook capture, comparing board, castling, en passant, clocks, history, hash and legal moves. Repair any mismatch and validate malformed FENs without panics.

### Fix playing correctness

- [ ] **C1 — Repair terminal and quiescence search.** Define one score perspective and mate distance convention; return immediately at game over; handle check evasions in quiescence. Add tactical FEN tests for mate in one, avoiding mate, stalemate and quiet evasions.
- [ ] **C2 — Repair transposition use.** Define remaining depth, exact/lower/upper bound semantics, mate score normalization and replacement policy. Add tests against search with TT disabled on transposing positions.
- [ ] **C3 — Complete game outcomes.** Track the halfmove clock and repeated positions through make/undo; implement the applicable draw conditions and test adjudication and search behavior.
- [ ] **C4 — Fix the HTTP game host.** Make reset visible to all handlers; reject invalid moves with one response; preserve exact special-move choice; serialize concurrent state changes. Cover the API with request-level tests.

### Build the developer instrument panel

- [ ] **D1 — Repair the existing UI.** Add promotion selection, correct board flip and asset handling, real evaluation values, visible loading/errors, and sequenced API updates. Confirm board, history and analysis stay aligned after move, undo and reset.
- [ ] **D2 — Expose bot analysis.** Add a bounded search request and response with depth, best move, score perspective, principal variation, elapsed time, nodes, NPS, cutoffs and TT statistics; make it visible in the UI with a board overlay. Keep calculation separate from applying a move.
- [ ] **D3 — Explain the backend.** Write a short engine walkthrough: square numbering, move encoding, `Next_move`, make/undo invariants, evaluation units, search recursion and TT meaning. Include one worked position from input to chosen move.

### Measure and grow strength

- [ ] **M1 — Create reproducible baselines.** Pin test positions and benchmark settings; capture speed and allocations for move generation, evaluation and fixed-depth search. Run paired version-versus-version games with saved PGNs and uncertainty estimates; do not infer Elo from NPS alone.
- [ ] **M2 — Add UCI and tournament play.** Implement a standalone UCI process with position setup, legal `bestmove`, depth/time controls and clean stop handling. Run repeatable matches against a pinned Stockfish build and at least one nearer-strength opponent under fixed conditions.
- [ ] **M3 — Improve search incrementally.** After C1/C2/M1, try iterative deepening, PV and TT move ordering, time management and selective search ideas one at a time. Require correctness checks and match results for each change.
- [ ] **M4 — Pursue a public rating.** Release a stable, documented UCI build; select a computer chess rating list and meet its current inclusion/testing conditions. Treat its published rating as distinct from local Stockfish match estimates.

## Learning recap

### Revisit this codebase before changing it

Follow one position through these files, writing down what each stage assumes and changes:

1. **Representation:** `src/chess_engine/notes.md`, `board/bitboard.go`, `board/board.go`, and `move_gen/moves.go`. Confirm square numbering, piece bitboards, the 16-bit move encoding, and the score units used for move ordering.
2. **Position lifecycle:** `gamestate/create_gs.go`, `fen_reader.go`, `intialise.go`, `make_move.go`, `undo.go`, and `zorbist.go`. Trace a normal move, a capture, a castle, en passant, and promotion; list every field that make/undo must restore.
3. **Legal moves:** `gamestate/gen_moves.go` and the pawn, king, check, pin, and magic move generators. See how attacks become legal moves, especially while in check or pinned. Compare the existing tests in `src/tests/test_move_gen` with the cases the generator can produce.
4. **Evaluation and search:** `src/chess_bot/evaluate.go`, `evaluate/`, `search.go`, `alpha_beta.go`, and `transposition.go`. Trace a two-ply search by hand: which side owns each score, where a cutoff occurs, and what a TT entry stores.
5. **Interfaces and measurements:** `src/server/`, `interface/src/Components/chess/DrawChess.jsx`, `board_ui.jsx`, `.api/api.jsx`, `src/tests/benchmark/`, and `perf/`. Understand how a move and analysis result reach the screen, then note what the current benchmarks actually measure. Read the UI files when working on D1/D2; they are not prerequisites for fixing the Go engine.

### Study beyond the codebase

- **For R2/R3/C1 now:** perft and divide as move-generation oracles; complete chess draw rules; negamax score perspective, mate distance, alpha-beta bounds, and why quiescence in check needs legal evasions.
- **For C2/M3:** transposition-table exact/lower/upper bounds, remaining-depth semantics, principal variation, iterative deepening, and move ordering. Study each when implementing it, not all before the first fix.
- **For M1/M2:** Go benchmarks and CPU/allocation profiling; UCI command flow and time controls; paired opening matches, Elo estimates, and uncertainty. Keep speed measurements and playing-strength evidence distinct.
- **For D1/D2:** React asynchronous state and request sequencing, plus an API schema that states whether an evaluation is static or searched and whose perspective it represents.

## AI workflow

The root `AGENTS.md` gives coding agents the project boundaries and verification rules. Take one work item at a time: reproduce the relevant failure, add a focused check, make the smallest understandable change, and report what was measured or could not be run. A custom skill would be useful later for repeatable perft or tournament runs, once those scripts and data formats exist; it would add little beyond the current guide today.

## UI decision

Keep React/Vite for now and repair its data flow. The developer-focused evaluation view is valuable, and the largest UI failures are local, fixable issues. UCI can connect the Go engine to existing chess GUIs for play and tournaments, while this UI remains the place for custom engine diagnostics. Reassess a rewrite only after D1/D2, using the cost of adding a desired diagnostic view as evidence.
