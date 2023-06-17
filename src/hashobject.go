package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func ExecuteHashObject(objectName string, isWrite bool) {
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
	n := fmt.Sprintf("%x", sha1.Sum(h))
	fmt.Println(n)

	if isWrite {
		if _, err := os.Stat(GIT_DIR_NAME + "objects/" + n[:2]); os.IsNotExist(err) {
			os.Mkdir(GIT_DIR_NAME+"objects/"+n[:2], 0777)
		}

		var c bytes.Buffer
		w := zlib.NewWriter(&c)
		w.Write(h)
		w.Close()

		g, err := os.Create(fmt.Sprintf(GIT_DIR_NAME+"objects/%s/%s", n[:2], n[2:]))
		if err != nil {
			panic(err)
		}
		defer g.Close()

		_, err = g.Write(c.Bytes())
		if err != nil {
			panic(err)
		}
	}
}
