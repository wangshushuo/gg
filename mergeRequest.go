package main

import (
  "bufio"
  "bytes"
  "fmt"
  "github.com/urfave/cli/v2"
  "log"
  "os"
  "os/exec"
  "strings"
)

func MergeRequest() cli.Command {
  sourceBranch := ReadBranchName()
  mergeRequest := cli.Command{
    Name:    "mr",
    Aliases: []string{"pr", "r"},
    Usage:   "post a merge request",
    Action: func(c *cli.Context) error {

      targetFlag := "-o merge_request.target=branch_name"
      createFlag := "-o merge_request.create"
      removeFlag := "-o merge_request.remove_source_branch"

      cmd := exec.Command("git", "push", "origin", "head:"+sourceBranch, targetFlag, createFlag, removeFlag)
      var out bytes.Buffer
      cmd.Stdout = &out
      fmt.Println(cmd.String())
      err := cmd.Run()
      if err != nil {
        log.Fatal(err)
      }
      fmt.Printf("in all caps: %q\n", out.String())
      return nil
    },
  }

  return mergeRequest
}

func ReadBranchName() string {
  fmt.Println("Enter a name:")
  reader := bufio.NewReader(os.Stdin)
  branchName, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }
  fmt.Print(branchName)
  return strings.TrimSpace(branchName)
}