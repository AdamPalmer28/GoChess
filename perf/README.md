# Performance tests

## Benchmark strategy

### Benches:

- Search:

  - Next_move [Implemented] - Sets up the gamestate ready for the next move (e.g. calc moves)

  - MoveSorting [Not designed & Not Implemented] - Sorts moves into an optimal order for the AI to evaluate

  - Evaluation [Implemented] - A full evaluation of the gamestate

  - FindBestMove [Implemented] - "ChessAI main", finds the best move for the AI to make

- Next_move [Implemented] - tests each stage of move generation

- MoveGen [Implemented]

- Evaluation - tests each section of the evaluation function
  - GetEvalMoveRays
  - EvalPieceCounts
  - PawnEval
  - KnightEval
  - BishopEval
  - RookEval
  - QueenEval
  - KingEval (not designed)
  - ... more

### Results
