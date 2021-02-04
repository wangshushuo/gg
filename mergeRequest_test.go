package main

import (
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
)

func TestGetNameOfCurrentBranch(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{"branches", "master"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNameOfCurrentBranch(); got != tt.want {
				t.Errorf("getNameOfCurrentBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeRequest(t *testing.T) {
	tests := []struct {
		name string
		want cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadBranchName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBranchName(); got != tt.want {
				t.Errorf("readBranchName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPickCurrentBranch(t *testing.T) {
	type args struct {
		branchesString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"pick current branch from list string of branch",
			args{
				branchesString: `
* master
111
aaa
`,
			},
			"master",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickCurrentBranch(tt.args.branchesString); got != tt.want {
				t.Errorf("pickCurrentBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}
