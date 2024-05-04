


printf "\nBenchmarking: chess engine...\n"

printf "\n/perf/benchmark_data/benchmark_$(date +"%Y-%m-%d").txt\n\n"



# benchmark and save results to perf/benchmark_[date].txt
printf "cwd: $(pwd) \n" 


go test -bench=. ./src/tests/... -run=^# -benchmem | tee perf/benchmark_data/benchmark_$(date +"%Y-%m-%d").txt