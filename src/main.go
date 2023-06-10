package main

import (
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var af = flag.String("f", "", "")
	flag.Parse()

	f, err := os.Open(*af)
	if err != nil {
		panic(fmt.Sprintf("Cannot open %s.", *af))
	}
	defer f.Close()

	r, err := zlib.NewReader(f)
	if err != nil {
		panic(fmt.Sprintf("Cannot read %s.", *af))
	}
	defer r.Close()

	io.Copy(os.Stdout, r)
}
