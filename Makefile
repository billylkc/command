linux: main.go
	go build -o bin/command main.go

windows: main.go
	GOOS=windows GOARCH=386 go build -o bin/command.exe main.go
