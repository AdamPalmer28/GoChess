cd src/


printf "\nTesting: chess_engine/board...\n"

go test ./chess_engine/board



printf "\nTesting: chess engine...\n"

go test ./tests/...
 