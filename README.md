use this command to build window version
"GOOS=windows GOARCH=386 go build -o zonai_windows main.go"

"GOOS=linux GOARCH=amd64 go build -o zonai_amd64 main.go"

dpkg --print-architecture
