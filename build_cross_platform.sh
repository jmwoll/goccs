GOOS=linux GOARCH=386 go build -o bin/goccs_linux_386
GOOS=linux GOARCH=amd64 go build -o bin/goccs_linux_amd64
GOOS=freebsd GOARCH=386 go build -o bin/goccs_freebsd_386
GOOS=windows GOARCH=386 go build -o bin/goccs_windows_386
GOOS=windows GOARCH=amd64 go build -o bin/goccs_windows_amd64
GOOS=darwin GOARCH=386 go build -o bin/goccs_darwin_386

