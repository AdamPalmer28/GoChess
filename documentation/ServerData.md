Data communicated between the go server (chess engine) and the frontend (interface)

# Gamedata

- gamestate
  - available moves
  - move list history
  - board
  - turn

# Bot data

Used to communicate with the bot to provide supporting information to the user

- AI move
  - move (and update gamestate)
  - evaluated score
  - search
    - depth
    - nodes (detail about types of nodes)
    - time
    - pv (principle variation)
    - best lines found
