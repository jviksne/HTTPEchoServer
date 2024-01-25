package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

// Config represents httpechoserver.yaml structure
type Config struct {
	Port         int    `yaml:"port"`
	CertFile     string `yaml:"certFile"`
	KeyFile      string `yaml:"keyFile"`
	ListenPath   string `yaml:"listenPath"`
	Log          bool   `yaml:"log"`
	MaxLogSizeMB int    `yaml:"maxLogSizeMB"`
}

var config Config

func echoHandler(w http.ResponseWriter, r *http.Request) {

	if config.Log {
		logRequest(r.RemoteAddr)
	}
	//fmt.Fprintf(w, "IP: %s\n\n", r.RemoteAddr)

	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

	fmt.Print(w, "\n\n")

	io.Copy(w, r.Body)
}

func logRequest(ip string) {

	logFilePath := getExeDirFilePath("httpechoserver.log")

	// Check if the log file size exceeds the maximum size
	if fileInfo, err := os.Stat(logFilePath); err == nil {
		sizeInMB := fileInfo.Size() / 1024 / 1024
		if int(sizeInMB) >= config.MaxLogSizeMB {
			return
		}
	}

	hasher := sha256.New()
	hasher.Write([]byte(ip))
	hashedIP := fmt.Sprintf("%x", hasher.Sum(nil))

	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer f.Close()

	timestamp := time.Now().Format(time.RFC3339)
	logEntry := fmt.Sprintf("%s - %s\n", timestamp, hashedIP)
	if _, err := f.WriteString(logEntry); err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

func getExeDir() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeDir := filepath.Dir(exePath)
	return exeDir
}

func getExeDirFilePath(fileName string) string {
	return filepath.Join(getExeDir(), fileName)
}

func main() {
	configData, err := os.ReadFile(getExeDirFilePath("httpechoserver.yaml"))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}

	if config.ListenPath == "" {
		config.ListenPath = "/echo"
	}

	http.HandleFunc(config.ListenPath, func(w http.ResponseWriter, r *http.Request) {
		echoHandler(w, r)
	})

	address := fmt.Sprintf(":%d", config.Port)
	fmt.Println("Started server on", address)
	if config.CertFile == "" && config.KeyFile == "" {
		err = http.ListenAndServe(address, nil)
	} else {
		err = http.ListenAndServeTLS(address, config.CertFile, config.KeyFile, nil)
	}
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
