package main

import (
	"os"
)

func main() {
	switch os.Args[1] {
	case "cat-file":
		ExecuteCatFile(os.Args[2])
	case "hash-object":
		ExecuteHashObject(os.Args[2], true)
	case "ls-files":
		ExecuteLsFiles()
	}
}
