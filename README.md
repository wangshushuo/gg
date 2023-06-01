https://github.com/wangshushuo/gg/actions/runs/run-id/artifacts/branch-name

# 命令行工具集合

注意事项：
- 使用本工具需要go环境
- 确保你的环境变量中有`$GOPATH/bin`
    - mac/linux: `echo -n 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc && source ~/.zshrc`
    - windows: (用户或系统)环境变量的path变量中加`%GOPATH%\bin`

## 安装

```
go install github.com/wangshushuo/gg@latest
GO111MODULE=on go install github.com/wangshushuo/gg@latest
```

## 命令：mr/pr/r

```shell script
gg mr
gg pr
gg r
```
向远程仓库（gitlab）发起一个 merge request 。目标分支是本地当前工作的分支。

输入命令后会要求输入一个临时分支名。

完成后，会将 merge request 的url复制的系统剪贴板。

### flag

```
gg mr -t feature-budget-2nd

gg mr -s bug-123
```
# 开发
更新代码后 `git tag v0.1.5` 更新版本号，`git push --tag` 上传版本号
