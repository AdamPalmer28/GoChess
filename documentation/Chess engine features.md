# Chess engine

## A record of all the features of the chess engine.

## Game state

---

## Move generation

---

## Search

- Alpha-beta pruning
  - Quiescence search
- Transposition table (needs to be improved)

---

## Evaluation

- Basic piece values
- Pawn Eval:
  - structure (doubled, isolated, backward, chain)
  - center control
  - promotion potential
- Knight Eval:
  - Basic middle sqs and edge sqs
- Bishop Eval:
  - move rays and xrays with HV pieces
- Rook Eval:
  - move rays and xrays with HV pieces
- Queen Eval:
  - move rays and xrays with HV pieces

---

## Testing

- Test MoveGen
- Make and Undo move
- Game state (check, checkmate, stalemate, draw)

### Benchmarking

---
