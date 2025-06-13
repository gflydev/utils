mod:
	go list -m --versions

critic:
	gocritic check -enableAll -disable=unnamedResult,unlabelStmt,hugeParam,singleCaseSwitch,builtinShadow,typeAssertChain ./...

security:
	gosec -exclude-dir=mysql,psql -exclude=G103,G115,G401,G501,G404 ./...

vulncheck:
	govulncheck ./...

lint:
	golangci-lint run ./...

test:
	go test -v -timeout 30s ./...

test.arr:
	go test -v -timeout 30s ./arr

test.coll:
	go test -v -timeout 30s ./coll

test.str:
	go test -v -timeout 30s ./str

test.seq:
	go test -v -timeout 30s ./seq

test.fn:
	go test -v -timeout 30s ./fn

test.num:
	go test -v -timeout 30s ./num

test.obj:
	go test -v -timeout 30s ./obj

all: critic security vulncheck lint test