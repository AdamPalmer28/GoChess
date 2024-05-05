printf "\nBenchmarking: chess engine...\n"

printf "\n/perf/benchmark_data/benchmark_$(date +"%Y-%m-%d").txt\n\n"


# benchmark and save results to perf/benchmark_[date].txt
go test ./src/tests/...  -bench=. -run=^# -benchmem | tee perf/benchmark_data/benchmark_$(date +"%Y-%m-%d").txt