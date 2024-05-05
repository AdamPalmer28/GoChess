---
# Chess to-do list

A ongoing to-do list of task for development of the chess engine.
---

### List

- Copy of gamestate

### Testing

- Test ChessBoard to uint64
  - Test Collisions
- Test TT is working correctly
- Copy of gamestate (create copy and check all of previous properties remain the same without affecting the original)

---

# Algorithm improvements

### High priority

- Search improvements
  - GetBestMove search before alpha-beta
  - move sorting
    - history heuristic
    - killer moves
    - take into account TT
  - iterative deepening
  - track line
  - TT improvements:
    - track line
    - track best line
- Evaluation:
  - King
    - king safety
      - castle rights
      - looming threats
      - defensive pieces
      - opponent activity

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
  - Board activity:
    - Identify weaknesses
      - Pawn structure
      - King safety
    - Identify threats
    - Identify critical squares
    - Identify outposts

---
