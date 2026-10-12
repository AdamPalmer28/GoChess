# =============================================================================
# DEV

install: go mod download
  cd interface && npm ci

# =============================================================================
# RUN

run: go run .

interface: cd interface && npm run dev

# =============================================================================
# TESTS

test: go test ./...

# =============================================================================
# PERFORMANCE

performance: go test ./src/tests/... -bench=. -run='^#' -benchmem

gamestate-benchmark: go test ./src/tests/... -bench='^(Benchmark_GenMoves)|(Benchmark_Next_move).*' -run='^#' -benchmem

ai-benchmark: go test ./src/tests/... -bench='^(Benchmark_Search_Time).*' -run='^#' -benchmem
