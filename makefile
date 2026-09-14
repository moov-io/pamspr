.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	go test ./... -short
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=60.0 time ./lint-project.sh
endif

.PHONY: bench
bench:
	go test ./pkg/pamspr -count=1 -run '^$$' -bench '^BenchmarkStreamingWriter_vs_TraditionalWriter$$' -benchmem | tee output.txt
