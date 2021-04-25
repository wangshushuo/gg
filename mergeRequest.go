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

	mergeRequest := cli.Command{
		Name:    "mr",
		Aliases: []string{"pr", "r"},
		Usage:   "发起一个 merge request",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "target", Aliases: []string{"t"}, Usage: "指定目标分支"},
			&cli.StringFlag{Name: "source", Aliases: []string{"s"}, Usage: "来源分支名"},
		},
		Action: func(c *cli.Context) error {
			flags := getFlag(c)
			s := flags["source"].(string)
			sourceBranch := readBranchName(s)

			var targetBranch string
			if t := flags["target"].(string); t != "" {
				targetBranch = t
				fmt.Printf("目标分支是【%s】\n", targetBranch)
			} else {
				targetBranch = getNameOfCurrentBranch()
				fmt.Printf("当前分支【%s】将做为目标分支。\n", targetBranch)
			}

			fmt.Println(" ")

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

func readBranchName(sourceBranch string) string {
	if sourceBranch == "" {
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
	fmt.Print("输入一个临时分支名：")
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

	return flagMap
}
