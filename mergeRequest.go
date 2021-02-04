package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
	"strings"
)

func MergeRequest() cli.Command {
	sourceBranch := readBranchName()
	targetBranch := getNameOfCurrentBranch()
	fmt.Printf("The name of current branch is [%s], and it will be the target branch.\n", targetBranch)
	fmt.Println(" ")
	mergeRequest := cli.Command{
		Name:    "mr",
		Aliases: []string{"pr", "r"},
		Usage:   "post a merge request",
		Action: func(c *cli.Context) error {

			targetFlag := "-o merge_request.target=" + targetBranch
			createFlag := "-o merge_request.create"
			removeFlag := "-o merge_request.remove_source_branch"

			cmd := exec.Command("git", "push", "origin", "head:"+sourceBranch, targetFlag, createFlag, removeFlag)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}

			gitMessage := string(out)
			fmt.Println(gitMessage)
			messages := strings.Fields(gitMessage)
			writeToClipboard(messages)
			return nil
		},
	}

	return mergeRequest
}

func readBranchName() string {
	fmt.Print("Enter a name as the temp source branch:")
	reader := bufio.NewReader(os.Stdin)
	branchName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(branchName)
}

func getNameOfCurrentBranch() string {
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

func pickCurrentBranch(branchesString string) string {
	currentBranchPrefix := "* "
	var branches []string
	branches = strings.Split(branchesString, "\n")
	indexOfCurrentBranch := 0
	for i, branch := range branches {
		isCurrent := strings.HasPrefix(branch, currentBranchPrefix)
		if isCurrent {
			indexOfCurrentBranch = i
			break
		}
	}
	currentBranchName := strings.TrimPrefix(branches[indexOfCurrentBranch], currentBranchPrefix)
	return currentBranchName
}

func writeToClipboard(messages []string) {
	for _, message := range messages {
		isMergeRequestURL := strings.Contains(message, "merge_requests")
		if isMergeRequestURL {
			err := clipboard.WriteAll(message)
			if err != nil {
				fmt.Println("Clipboard operation failed 😫")
				log.Fatal(err)
			}
			fmt.Println("The merge_request_url has been written to the operation clip 😎")
			break
		}
	}
}
