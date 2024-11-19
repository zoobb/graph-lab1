EXECUTABLE := graph_lab1.exe

SOURCE := main.go

WINDOWS_FOLDER := /mnt/c/Users/hephaestusWin/Documents/uni/graph/lab1

build:
	GOOS=windows GOARCH=amd64 go build -ldflags "-H=windowsgui" -o $(WINDOWS_FOLDER)/$(EXECUTABLE) $(SOURCE)

clean:
	rm -f $(WINDOWS_FOLDER)/$(EXECUTABLE)

rebuild: clean build

install-deps:
	go get -u ./...

.PHONY: build clean rebuild install-deps