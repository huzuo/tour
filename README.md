# tour

> 本项目基于第三方开源库Cobra和标准库strings、unicode快速搭建一个CLI应用程序，实现了多种模式的单词转换功能，此功能在日常工作中较为实用，因为我们经常要对输入、输出数据进行各种类型的转换和拼装。

项目目录结构

```shell
tour
├── cmd
│   ├── root.go	#放置根命令
│   └── word.go	#放置单词转化的子命令word
├── go.mod
├── go.sum
├── internal
│   └── word
│       └── word.go	#转换逻辑代码
└── main.go
```

使用示例：

```shell
#将项目clone到本地
#查看帮助文档
[root@kvm /usr/local/go/code/src/tour]# go run main.go word -h
该子命令支持各种单词格式转换，模式如下：
1：全部单词转为大写
2：全部单词转为小写
3：下划线单词转为大写驼峰单词
4：下划线单词转为小写驼峰单词
5：驼峰单词转为下划线单词

Usage:
   word [flags]

Flags:
  -h, --help         help for word
  -m, --mode int8    请输入单词转换模式
  -s, --str string   请输入单词内容
#全部单词转为大写
[root@kvm /usr/local/go/code/src/tour]# go run main.go word -s=hello_word -m=1
2022/09/27 23:17:36 输出结果：HELLO_WORD
#全部单词转为小写
[root@kvm /usr/local/go/code/src/tour]# go run main.go word -s=hELlo_woRd -m=2
2022/09/27 23:18:07 输出结果：hello_word
#下划线单词转为大写驼峰单词
[root@kvm /usr/local/go/code/src/tour]# go run main.go word -s=hello_word -m=3
2022/09/27 23:18:24 输出结果：HelloWord
#下划线单词转为小写驼峰单词
[root@kvm /usr/local/go/code/src/tour]# go run main.go word -s=Hello_word -m=4
2022/09/27 23:18:37 输出结果：helloWord
#驼峰单词转为下划线单词
[root@kvm /usr/local/go/code/src/tour]# go run main.go word -s=HelloWord -m=5
2022/09/27 23:18:53 输出结果：hello_word
```

