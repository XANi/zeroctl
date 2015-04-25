all: dep
	go fmt
	gom exec go build

dep:
	gom install

test:
	gom exec go test
