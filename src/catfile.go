package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetGitObjectType(b []byte) string {
	buff := []byte{}
	for _, c := range b {
		if c == 0 {
			break
		}
		buff = append(buff, c)
	}
	return strings.Split(string(buff), " ")[0]
}

func PrintTree(c []byte) {
	buff := []byte{}
	isId := false
	for _, c := range c {
		buff = append(buff, c)
		if !isId && c == 0 {
			fmt.Print(string(buff))
			buff = []byte{}
			isId = true
		} else if isId && len(fmt.Sprintf("%x", buff)) == 40 {
			fmt.Printf("%x\n", buff)
			buff = []byte{}
			isId = false
		}
	}
}

func ExecuteCatFile(gitObjectId string) {
	p := GetFilePathFromGitObjectId(gitObjectId)

	f, err := os.Open(p)
	if err != nil {
		panic(fmt.Sprintf("Cannot open %s.", p))
	}
	defer f.Close()

	r, err := zlib.NewReader(f)
	if err != nil {
		panic(fmt.Sprintf("Cannot extract git object %s.", gitObjectId))
	}
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		panic(fmt.Sprintf("Cannot load git object %s.", gitObjectId))
	}
	defer r.Close()

	c := bytes.SplitAfterN(b, []byte{0}, 2)[1]

	switch GetGitObjectType(b) {
	case "commit":
		fmt.Println(string(c))
	case "tree":
		PrintTree(c)
	case "blob":
		fmt.Println(string(c))
	}
}
