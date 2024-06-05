
package main

import (
    "testing"
)

func TestGenerateChallenge(t *testing.T) {
    challenge := generateChallenge(20)
    if len(challenge) != 20 {
        t.Errorf("Expected challenge length of 20, but got %d", len(challenge))
    }
}

func TestCheckProofOfWork(t *testing.T) {
    challenge := "testchallenge"
    nonce := "0000"
    if !checkProofOfWork(challenge, nonce, 4) {
        t.Errorf("Expected valid proof of work")
    }
}
