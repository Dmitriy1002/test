
package main

import (
    "testing"
)

func TestFindNonce(t *testing.T) {
    challenge := "testchallenge"
    nonce := findNonce(challenge, 4)
    if !checkProofOfWork(challenge, nonce, 4) {
        t.Errorf("Expected valid proof of work")
    }
}
