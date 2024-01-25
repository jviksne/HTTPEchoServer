# HTTP Echo Server

Simple Go based HTTP(s) server echoing all the headers and message body of the request to
in the body of the response. 

## How to run

### 1. Select the Precompiled Binary for Your OS:
- `httpechoserver` for Linux (AMD64)
- `httpechoserver.exe` for Windows (AMD64)
### 2. Configure the Server:
  Edit the parameters stored in httpechoserver.yaml. Place it in the same directory as the binary executable.
### 3. Set Up the Log File:
  Place httpechoserver.log in the same directory as the binary executable and ensure it is writable by the user running the executable.
### 4. Launch the Executable:
  Run the executable file.
### 5. To Enable a Linux Service (e.g., on Ubuntu):
#### 5.1 Edit and copy the `httpechoserver.service` file to
  `/etc/systemd/system/httpechoserver.service`.
#### 5.2 Enable the service using:
```
sudo systemctl enable httpechoserver
```
#### 5.3 Control the service:
```
# Start the service
sudo systemctl start httpechoserver

# Check status
sudo systemctl status httpechoserver

# Restart the service
sudo systemctl restart httpechoserver

# Stop the service
sudo systemctl stop httpechoserver

# Disable the service
sudo systemctl disable httpechoserver
```

## How to rebuild

### 1. Install Go
   Make sure Go is installed on your system.
### 2. Download the Repository:
   Download and extract the repository.
### 3. Download Dependencies:
   From the command line, run:
```
go mod download
```
### 4. Build Instructions:
  - For Windows:
```
SET GOOS=windows
SET GOARCH=amd64
go build -o httpechoserver.exe
```
  - For Linux:
```
SET GOOS=linux
SET GOARCH=amd64
go build -o httpechoserver
```
