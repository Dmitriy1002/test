package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

const (
	difficulty = 4
	serverAddr = "localhost:5000"
)

func findNonce(challenge string, difficulty int) string {
	var nonce string
	for i := 0; ; i++ {
		nonce = fmt.Sprintf("%x", i)
		data := challenge + nonce
		hash := sha256.Sum256([]byte(data))
		hashStr := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashStr, strings.Repeat("0", difficulty)) {
			break
		}
	}
	return nonce
}

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	challenge := strings.TrimSpace(string(buf[:n]))
	fmt.Println("Received challenge:", challenge)

	nonce := findNonce(challenge, difficulty)
	conn.Write([]byte(nonce + "\n"))

	n, _ = conn.Read(buf)
	response := strings.TrimSpace(string(buf[:n]))
	fmt.Println("Server response:", response)
}
