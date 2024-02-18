---
# Chess project to-do list

A ongoing to-do list of task for development of the chess project.
---

## High priority

- Options menu (left of the board)

  - New game
  - Undo
  - Load game (png)
  - AI eval
  - ...

- Eval bar (on right of board)

### Interface

- Footer gamestate UI
  - turn
- Footer Settings UI (with settings)
- Footer AI UI

  - Calculate next move

- Promotion UI (popup)

- Basic display of Move history log (side UI game tab)
- Analysis section (side UI)

  - select areas of the program like gamestate, ai, ... (expandable)
    - cause the board to high light areas of interest

- Flexible board square drawing

### Backend server

- Create exports to packageChessData for gamestate

  - bitboards

- Undo end point

- Calculate AI move
  (take in setting parameters)

### Misc tasks

- Create Makefile of commands
- Docker container for chess engine
- Docker container for chess program (engine + interface)
