---
# Chess to-do list

A ongoing to-do list of task for development of the chess engine.
---

### High priority

- add testing for all 'test to-dos'

### General

- gamestate / move gen
  - move gen for incheck
    idea: gen the move rays for the threat squares (enemy piece and attack rays between enemy piece and king) plus king moves
    - move gen for incheck
      - pawn gen - need to improve
- piece pins:
  idea: make if piece on king safety ray and moves must be on ray

### Low priority

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
