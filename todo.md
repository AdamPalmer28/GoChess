---
# Chess to-do list

A ongoing to-do list of task for development of the chess engine.
---

### High priority

- undo testing
- convert best move num to text

### General

- Evaluation:
  - General
    - activty
    - defenders
    - rays pieces:
      - avaliable movement
      - attacking movement
  - Pawn
    - promotion
    - centre control
    - chain
  - Knight
    - middle controls
    - attacking sqs
    - outposts
  - Bishop
    - xrays
  - Rook
    - xrays
    - open files
  - Queen
    - xrays
  - King
    - king safety
      - castle rights
      - looming threats
      - defensive pieces
      - opponent activity

### Low priority

- ChessBoard to uint64
- Pawn ray generation pre-computed
  - pre-computed [2][2][64]board.Bitboard
    [move/capture][color][square]

---

# Ideas

---

- Test environment for testing the engine against previous versions
- Test scripts - Correct moves i.e. as expected - Function outputs - Benchmark test & function's speed

# Go Notes:

packages:
lower case = private
upper case = public

&x is the address of x
\*x is the value of x
