package main

// Solution for both questions 7 and 8 from quiz.
// Final xor could cancel out things, we take this advantage.

import (
	"crypto/aes"
	"fmt"
	"encoding/hex"
	"bytes"
)

const size = 16

func aesDec(k, m []byte) []byte {
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(err)
	}

	dst := make([]byte, len(m))
	block.Decrypt(dst, m)

	return dst
}

func aesEnc(k, m []byte) []byte {
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(err)
	}

	dst := make([]byte, len(m))
	block.Encrypt(dst, m)

	return dst
}

func f1(x, y []byte) []byte {
	dst := aesEnc(y, x)

	for i, b := range dst {
		dst[i] = b ^ y[i]
	}

	return dst
}

func f2(x, y []byte) []byte {
	dst := aesEnc(x, x)

	for i, b := range dst {
		dst[i] = b ^ y[i]
	}

	return dst
}

func solveF1() {
	// f1(x, y) = aes.enc(x, y) xor y
	// "xor y" should cancel out aes to all zeros

	zeros := make([]byte, size)

	y1 := zeros
	x1 := aesDec(y1, zeros)

	ones := make([]byte, size)
	for i, _ := range ones {
		ones[i] = 255
	}

	y2 := ones
	x2 := aesDec(y2, ones)

	if !bytes.Equal(f1(x1, y1), f1(x2, y2)) {
		panic("incorrect solution for f1")
	}

	fmt.Println("f1 solution:")
	fmt.Println("x1", hex.EncodeToString(x1))
	fmt.Println("y1", hex.EncodeToString(y1))
	fmt.Println("x2", hex.EncodeToString(x2))
	fmt.Println("y2", hex.EncodeToString(y2))
}

func solveF2() {
	// f2(x, y) = aes.enc(x, x) xor y
	// "xor y" should cancel out aes to all zeros

	zeros := make([]byte, size)

	x1 := zeros
	y1 := aesEnc(x1, x1)

	ones := make([]byte, size)
	for i, _ := range ones {
		ones[i] = 255
	}

	x2 := ones
	y2 := aesEnc(x2, x2)

	if !bytes.Equal(f2(x1, y1), f2(x2, y2)) {
		panic("incorrect solution for f2")
	}

	fmt.Println("f2 solution:")
	fmt.Println("x1", hex.EncodeToString(x1))
	fmt.Println("y1", hex.EncodeToString(y1))
	fmt.Println("x2", hex.EncodeToString(x2))
	fmt.Println("y2", hex.EncodeToString(y2))
}

func main() {
	solveF1()
	solveF2()
}

