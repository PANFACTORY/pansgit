package main

import (
	"fmt"
	"os"
)

func GetFilePathFromGitObjectId(s string) string {
	return fmt.Sprintf("./.git/objects/%s/%s", string([]rune(s)[:2]), string([]rune(s)[2:]))
}

func main() {
	switch os.Args[1] {
	case "cat-file":
		ExecuteCatFile(os.Args[2])
	}
}
