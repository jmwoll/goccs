GOOS=linux GOARCH=386 go build -o bin/linux_386
GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64
GOOS=freebsd GOARCH=386 go build -o bin/freebsd_386
GOOS=windows GOARCH=386 go build -o bin/windows_386
GOOS=windows GOARCH=amd64 go build -o bin/windows_amd64
GOOS=darwin GOARCH=386 go build -o bin/darwin_386

