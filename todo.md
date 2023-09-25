---
# Chess to-do list

A ongoing to-do list of task for development of the chess engine.
---

### High priority

- undo testing

  - moves
    - normal move
    - capture
    - enpassant
    - promotion
    - castle
      - castle rights
  - history lists
  - move gen
    - incheck

- convert best move num to text
  - detect mate

### General

- Evaluation:
  - General
    - activty
    - defenders
    - critical squads
  - Knight
    - outposts
    - activity
      - multiple attacks
  - King
    - king safety
      - castle rights
      - looming threats
      - defensive pieces
      - opponent activity

### Low priority

- ChessBoard to uint64

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
