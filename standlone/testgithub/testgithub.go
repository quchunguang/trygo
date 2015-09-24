package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"os/exec"
	"strings"
)

func main() {
	client := github.NewClient(nil)
	repos, _, err := client.Repositories.List("quchunguang", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, r := range repos {
		s := *r.CloneURL
		s = strings.TrimPrefix(s, "https://")
		s = strings.TrimSuffix(s, ".git")
		fmt.Println(s)
	}

	s := "github.com/quchunguang/projecteuler"
	_, err = exec.Command("go", "get", s).Output()
	if err != nil {
		fmt.Println(err)
	}
}
