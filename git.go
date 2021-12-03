package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func tempBranchName() string {
	return "temp_" + strconv.FormatInt(time.Now().UnixNano(), 10)
}

func lastCommitMessage() string {
	//	git log -1 --pretty=%B
	cmd := exec.Command("git", "log", "-1", "--pretty=%B")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	gitMessage := string(out)
	fmt.Println(gitMessage)
	return strings.TrimSpace(gitMessage)
}
func nameOfCurrentBranch() string {
	cmd := exec.Command("git", "branch")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	branch := pickCurrentBranch(out.String())
	return branch
}

func fetchMergeRequest(source, target string) string {
	targetFlag := "-o merge_request.target=" + target
	createFlag := "-o merge_request.create"
	removeFlag := "-o merge_request.remove_source_branch"

	cmd := exec.Command("git", "push", "origin", "head:"+source, targetFlag, createFlag, removeFlag)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	gitMessage := string(out)
	return gitMessage
}
