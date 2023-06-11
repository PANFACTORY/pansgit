package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type indexEntry struct {
	cSeconds       uint32
	cNanosecond    uint32
	mSeconds       uint32
	mNanosecond    uint32
	objectType     string
	unixPermission string
	fileSize       uint32
	id             string
	stage          string
	nameLength     uint16
	name           string
}

func getIndexEntries(b []byte, n int) []indexEntry {
	e := []indexEntry{}
	i := 0
	for j := 0; j < n; j++ {
		l := binary.BigEndian.Uint16(b[i+60:i+62]) & 0xfff
		e = append(e, indexEntry{
			binary.BigEndian.Uint32(b[i : i+4]),
			binary.BigEndian.Uint32(b[i+4 : i+8]),
			binary.BigEndian.Uint32(b[i+8 : i+12]),
			binary.BigEndian.Uint32(b[i+12 : i+16]),
			fmt.Sprintf("%b", binary.BigEndian.Uint16(b[i+26:i+28])>>12),
			fmt.Sprintf("%o", binary.BigEndian.Uint16(b[i+26:i+28])&0x1ff),
			binary.BigEndian.Uint32(b[i+36 : i+40]),
			fmt.Sprintf("%x", b[i+40:i+60]),
			fmt.Sprintf("%b", (binary.BigEndian.Uint16(b[i+60:i+62])&0x3000)>>12),
			l,
			string(b[i+62 : i+62+int(l)]),
		})
		i += 62 + int(l) + 8 - (i+62+int(l))%8
	}
	return e
}

func ExecuteLsFiles() {
	f, err := os.Open(".git/index")
	if err != nil {
		panic("Cannot open index file.")
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic("Cannot read index file.")
	}

	i := int(binary.BigEndian.Uint32(b[8:12]))
	c := b[12:]

	for _, e := range getIndexEntries(c, i) {
		fmt.Printf("%s%s %s %s\t%s\n", e.objectType[:3], e.unixPermission, e.id, e.stage, e.name)
	}
}
