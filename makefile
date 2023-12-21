# Make docker container for chess engine


# -----------------------------------------------------------------------------
# Chess engine

ENGINE_PATH = src

run: 
	ls
	go run $(ENGINE_PATH)/main.go

benchmarking: 
	cd src/ 
	printf "\nBenchmarking: chess engine...\n"
	go test -bench=. ./tests/...

test:
	cd src/ 
	printf "\nTesting: chess_engine/board...\n"
	go test ./chess_engine/board
	printf "\nTesting: chess engine...\n"
	go test ./tests/...

