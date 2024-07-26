# go-cutter
`go-cutter`是一个命令行工具，用于快速使用模板或克隆现有结构来脚手架Go项目。

# 功能特性
- 在模板项目根路径下执行命令可创建新的Go项目
- 自动替换 import 路径
- 自动更新 go.mod 文件中的模块名称
- 自动删除 .git 目录

***注意：一定要在模板项目的根路径下执行命令***
# 安装
```shell
go install github.com/morehao/go-cutter@latest
```
# 使用方法
## 初始化新项目
```shell
cd /appTemplatePath
go-cutter -d /yourAppPath
```
- `-d, --destination`：新项目的目标目录，例如：`/user/myApp`。此参数为必填项。


