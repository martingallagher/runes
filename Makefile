.PHONY: bench

bench:
	@go test -v -bench=. -tags=std > old.txt
	@go test -v -bench=. > new.txt
	@benchcmp old.txt new.txt
	@rm -f old.txt new.txt
