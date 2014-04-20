package main

// You should pick solution with distinguishable pattern

import (
	"encoding/hex"
	"fmt"
)

func check(lh, rh string) {
	l, err := hex.DecodeString(lh)
	if err != nil {
		panic(err)
	}

	r, err := hex.DecodeString(rh)
	if err != nil {
		panic(err)
	}

	xor := make([]byte, len(l))
	for i, _ := range l {
		xor[i] = l[i] ^ r[i]
	}

	fmt.Printf("%s %s\n", lh, rh)
	fmt.Printf("  %v\n", xor)
}

func main() {
	check("9f970f4e932330e4", "6068f0b1b645c008")
	check("7b50baab07640c3d", "ac343a22cea46d60")
	check("9d1a4f78cb28d863", "75e5e3ea773ec3e6")
	check("4af532671351e2e1", "87a40cfa8dd39154")
}
