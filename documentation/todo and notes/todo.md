---
# Chess project to-do list

A ongoing to-do list of task for development of the chess project.
---

## High priority

### Interface

- Options menu (left of the board)

  - New game
  - Undo
  - Load game (png)
  - AI eval
  - ...

- Eval bar (on right of board)
- Footer UI

  - finish gamestate UI
  - Settings UI (with settings)
  - AI UI
    - Calculate next move

- Promotion UI (popup)

- Side UI

  - Basic display of Move history log (game tab)
  - Analysis section:
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
