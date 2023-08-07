package main

import (
	"bytes"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
	"log"
	"os/exec"
	"strings"
)

func MergeRequest() cli.Command {

	mergeRequest := cli.Command{
		Name:    "mr",
		Aliases: []string{"pr", "r"},
		Usage:   "发起一个 merge request",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "target", Aliases: []string{"t"}, Usage: "指定目标分支"},
			&cli.StringFlag{Name: "source", Aliases: []string{"s"}, Usage: "来源分支名"},
			&cli.StringFlag{Name: "assign", Aliases: []string{"a"}, Usage: "指派"},
		},
		Action: func(c *cli.Context) error {
			var sourceBranch = c.String("target")
			var targetBranch = c.String("source")
			var assign = c.String("assign")

			if sourceBranch == "" {
				sourceBranch = tempBranchName()
			}

			if targetBranch == "" {
				targetBranch = getNameOfCurrentBranch()
			}

			fmt.Println("发起MR")
			fmt.Println("临时分支: " + sourceBranch)
			fmt.Println("目标分支: " + targetBranch)
			fmt.Println("目标分支: " + targetBranch)
			fmt.Println("指    派: " + assign)

			gitMessage := fetchMergeRequest(sourceBranch, targetBranch, assign)
			//gitMessage := "1"
			fmt.Println(gitMessage)
			messages := strings.Fields(gitMessage)
			writeToClipboard(messages)
			return nil
		},
	}

	return mergeRequest
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
				fmt.Println("Clipboard 操作失败 😫")
				log.Fatal(err)
			}
			fmt.Println("The merge_request_url 已经添加到 Clipboard 可以直接 ctrl + v 了 😎")
			break
		}
	}
}

func getFlag(c *cli.Context) map[string]interface{} {
	flagMap := make(map[string]interface{})
	flagMap["target"] = c.String("target")
	flagMap["source"] = c.String("source")
	flagMap["assign"] = c.String("assign")

	return flagMap
}
