## 命令行工具集合

主命令：gg

### 子命令：mr/pr/r

```shell script
gg r
```
向远程仓库（gitlab）发起一个 merge request 。目标分支是本地当前工作的分支。

输入命令后会要求输入一个临时分支名。

完成后，会将 merge request 的url复制的系统剪贴板。