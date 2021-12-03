package main

import (
	"reflect"
	"testing"
)

func Test_branchName1(t *testing.T) {
	type args struct {
		lastCommitMsg string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "提取bug及编码",
			args: args{
				lastCommitMsg: "笑果-bug-0000-自动滚动到当前时间或者开始时间；修复时间问题；修复遮罩问题；",
			},
			want: []string{"bug", "0000"},
		},
		{
			name: "提取bug及编码",
			args: args{
				lastCommitMsg: "笑果-feature-12345-自动滚动到当前时间或者开始时间；修复时间问题；修复遮罩问题；",
			},
			want: []string{"feature", "12345"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := branchName(tt.args.lastCommitMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("branchName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("branchName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
