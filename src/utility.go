package main

import "fmt"

var GIT_DIR_NAME = ".git_dummy/"

func GetFilePathFromGitObjectId(s string) string {
	return fmt.Sprintf(GIT_DIR_NAME+"objects/%s/%s", string([]rune(s)[:2]), string([]rune(s)[2:]))
}
