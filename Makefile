linux: main.go
	go build -o cmd/command main.go

windows: main.go
	GOOS=windows GOARCH=386 go build -o cmd/command.exe main.go
