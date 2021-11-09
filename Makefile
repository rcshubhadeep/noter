build: setup
	go build -o build/ github.com/rcshubhadeep/noter/
	
clean:
	rm -rf tmp build

setup:
	mkdir -p tmp build/linux-amd64 build/windows-amd64 build/macos-amd64

linux:
	GOOS=linux GOARCH=amd64 go build -o build/linux-amd64/ github.com/rcshubhadeep/noter/...

windows:
	GOOS=windows GOARCH=amd64 go build -o build/win-amd64/ github.com/rcshubhadeep/noter/...

darwin:
	GOOS=darwin GOARCH=amd64 go build -o build/macos-amd64/ github.com/rcshubhadeep/noter/...

all: clean linux windows darwin

show_arch: 
	go tool dist list -json