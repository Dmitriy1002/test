package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

const (
	difficulty = 4
	port       = ":5000"
)

var quotes = []string{
	"Quote 1: Wisdom is the reward for surviving our own stupidity.",
	"Quote 2: The only true wisdom is in knowing you know nothing.",
	"Quote 3: The unexamined life is not worth living.",
}

func generateChallenge(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func checkProofOfWork(challenge, nonce string, difficulty int) bool {
	data := challenge + nonce
	hash := sha256.Sum256([]byte(data))
	hashStr := hex.EncodeToString(hash[:])
	return strings.HasPrefix(hashStr, strings.Repeat("0", difficulty))
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	rand.Seed(time.Now().UnixNano())

	challenge := generateChallenge(20)
	conn.Write([]byte(fmt.Sprintf("POW challenge: %s
", challenge)))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	nonce := strings.TrimSpace(string(buf[:n]))

	if checkProofOfWork(challenge, nonce, difficulty) {
		quote := quotes[rand.Intn(len(quotes))]
		conn.Write([]byte(fmt.Sprintf("Success! Here's your quote: %s
", quote)))
	} else {
		conn.Write([]byte("Invalid proof of work
"))
	}
}

func main() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server is listening on port", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
