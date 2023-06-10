package main

import (
	"compress/zlib"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(fmt.Sprintf("Cannot read %s.", *af))
	}
	defer r.Close()

	s := strings.SplitAfterN(string(b), string([]byte{0}), 2)

	fmt.Printf("%s", s)
}
