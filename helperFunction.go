package main

import (
	"crypto/rand"
	"fmt"
	"log"
	mrand "math/rand" // Aliases during import to avoid naming conflicts.
)

func GenerateRandId() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	id := fmt.Sprintf("%X-%X", b[0:4], b[4:8])
	return id
}

func GenRandomNumForRange(numRange int) int {
	return mrand.Intn(numRange)
}
