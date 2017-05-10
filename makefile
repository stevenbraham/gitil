all: build
build: 
	rm -rf out
	mkdir out
	mkdir out/windows
	mkdir out/mac
	mkdir out/linux
	GOOS=windows GOARCH=amd64 go build -o out/windows/gitil-64.exe
	GOOS=windows GOARCH=386 go build -o out/windows/gitil-32.exe
	GOOS=darwin GOARCH=amd64 go build -o out/mac/gitil-64
	GOOS=darwin GOARCH=386 go build -o out/mac/gitil-32
	GOOS=linux GOARCH=amd64 go build -o out/linux/gitil-64
	GOOS=linux GOARCH=386 go build -o out/linux/gitil-32
