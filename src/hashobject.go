package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func ExecuteHashObject(objectName string) {
	f, err := os.Open(objectName)
	if err != nil {
		panic(fmt.Sprintf("Cannot open %s.", objectName))
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(fmt.Sprintf("Cannot read %s.", objectName))
	}

	s, err := f.Stat()
	if err != nil {
		panic(fmt.Sprintf("Cannot read propaty of %s", objectName))
	}

	h := append(append([]byte(fmt.Sprintf("blob %d", s.Size())), []byte{0}...), b...)
	fmt.Printf("%x", sha1.Sum(h))
}
