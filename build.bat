echo Compiling for Windows
SET GOOS=windows
SET GOARCH=amd64
go build -o httpechoserver.exe

echo Compiling for Linux
SET GOOS=linux
SET GOARCH=amd64
go build -o httpechoserver