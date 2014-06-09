package main

import (
    "fmt"
    "math/rand"
)

// Taken from github.com/revel/revel/revel/new.go

const alphaNumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func generateSecret() string {
    chars := make([]byte, 64)
    for i := 0; i < 64; i++ {
        chars[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
    }
    return string(chars)
}

func main() {
    fmt.Println(generateSecret())
}