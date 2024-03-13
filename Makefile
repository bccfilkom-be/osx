.PHONY: test
test:
	go test -v -race -buildvcs

.PHONY: bench
bench:
	go test -bench .
