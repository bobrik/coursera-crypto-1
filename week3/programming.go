package main

//

import (
	"fmt"
	"encoding/hex"
	"os"
	"crypto/sha256"
)

const size = 1024

func hash(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	s, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("size: %d, blocks: %d\n", s.Size(), s.Size() / size)

	b := s.Size() / size
	if b * size < s.Size() {
		b += 1
	}

	a := []byte{}

	for i := b - 1; i >= 0; i-- {
		l := int64(size)
		if i == b - 1 {
			l = s.Size() - (b - 1) * size
		}

		buf := make([]byte, l)
		f.ReadAt(buf, i * size)

		h := sha256.New()
		h.Write(append(buf, a...))
		a = h.Sum([]byte{})
	}

	return a
}

func main() {
	fmt.Println("test hash is:", hex.EncodeToString(hash("./6 - 2 - Generic birthday attack (16 min).mp4")))
	fmt.Println("target hash is:", hex.EncodeToString(hash("./6 - 1 - Introduction (11 min).mp4")))
}
