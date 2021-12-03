package main

import (
	"errors"
	"regexp"
)

func branchName(lastCommitMsg string) ([]string, error) {
	reg, _ := regexp.Compile(`(bug)|(feature)-(\d+)-`)
	sub := reg.FindStringSubmatch(lastCommitMsg)
	if len(sub) >= 3 {
		return sub[1:], nil
	}
	return []string{}, errors.New("名称中不包含bug")
}
