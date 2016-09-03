./cmd -profile
go tool pprof  ./cmd cpu.pprof

go tool pprof --pdf ./cmd cpu.pprof  > cgraph.pdf