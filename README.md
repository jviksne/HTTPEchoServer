# HTTP Echo Server

Simple Go based HTTP(s) server echoing all the headers and message body of the request to
in the body of the response. 

## How to run

1. Select the precompiled binary for your OS:
- httpechoserver: Linux (AMD64)
- httpechoserver: Windows (AMD64)
2. Edit the paremeters stored in httpechoserver.yaml. Place in the same directory as the binary executable.
3. Place httpechoserver.log in the same directory as the binary executable and ensure it is writable for the user that will be running the executable.
4. Launch the executable.
5. To enable a Linux service on, e.g., Ubuntu:
5.1 Edit and copy the httpechoserver.service to: /etc/systemd/system/httpechoserver.service
5.2 Enable the service:
```
sudo systemctl enable httpechoserver
```
5.3 Control the service:
```
# Start the service
sudo systemctl start httpechoserver

# Check status
sudo systemctl status httpechoserver

# Restart the service
sudo systemctl restart httpechoserver
```

## How to rebuild

1. Install Go
2. Download the repository and extract to some directory.
3. From CMD run:
```
go mod download
```
4. To build for Windows run:
```
SET GOOS=windows
SET GOARCH=amd64
go build -o httpechoserver.exe
```

To build for Linux run:
```
SET GOOS=linux
SET GOARCH=amd64
go build -o httpechoserver
```
