GO_VERSION:=$(shell go version)

fmt:
	# TODO: When Go 1.9 is released vendor folder should be ignored automatically
	bash -c 'go list ./... | grep -v vendor | xargs -n1 go fmt'

lintall:
	bash -c 'gometalinter --tests --vendor --deadline=300s ./...'

lint:
	bash -c 'gometalinter --disable-all -E vet -E gofmt -E misspell -E ineffassign -E goimports -E deadcode -E gocyclo --tests --vendor ./...'

golint:
	# TODO: When Go 1.9 is released vendor folder should be ignored automatically
	bash -c 'go list ./... | grep -v vendor | grep -v mocks | xargs -n1 golint'

test:
	# TODO: When Go 1.9 is released vendor folder should be ignored automatically
	bash -c 'go list ./... | grep -v vendor | xargs -n1 go test -v -timeout=60s -cover'

test-with-coverage:
	bash -c 'go test -v -race -coverprofile=coverage.out -covermode=atomic .'

cover:
	bash -c 'go tool cover -func=coverage.out && go tool cover -html=coverage.out'

bench:
	# TODO: When Go 1.9 is released vendor folder should be ignored automatically
	bash -c 'go list ./... | grep -v vendor | xargs -n1 go test -count=5 -run=NONE -bench . -benchmem'

profile:
	mkdir pprof
	go test -count=10 -run=NONE -bench . -benchmem -o pprof/test.bin -cpuprofile pprof/cpu.out -memprofile pprof/mem.out
	go tool pprof --svg pprof/test.bin pprof/mem.out > mem.svg
	go tool pprof --svg pprof/test.bin pprof/cpu.out > cpu.svg
	rm -rf pprof
	go test -count=10 -run=NONE -bench=BenchmarkGlg -benchmem -o pprof/test.bin -cpuprofile pprof/cpu-glg.out -memprofile pprof/mem-glg.out
	go tool pprof --svg pprof/test.bin pprof/cpu-glg.out > cpu-glg.svg
	go tool pprof --svg pprof/test.bin pprof/mem-glg.out > mem-glg.svg
	rm -rf pprof
	go test -count=10 -run=NONE -bench=BenchmarkDefaultLog -benchmem -o pprof/test.bin -cpuprofile pprof/cpu-def.out -memprofile pprof/mem-def.out
	go tool pprof --svg pprof/test.bin pprof/mem-def.out > mem-def.svg
	go tool pprof --svg pprof/test.bin pprof/cpu-def.out > cpu-def.svg
	rm -rf pprof
	mv ./*.svg bench/
