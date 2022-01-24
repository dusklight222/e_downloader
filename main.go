package main

import (
	"fmt"
	"os"

	// "fmt"
	// "os"
	"strings"
)

var (
	// 用于更换的域名头
	DomainHeader string = "1252524126.vod2.myqcloud.com"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("请输入您的URL文件")
		return
	}
	path := list[1]
	// root, _ := os.Getwd()
	// 获取文件名
	num1 := strings.LastIndex(path, ".txt")
	// 获取文件名，以此作为根目录
	root := path[:num1]
	fmt.Println(root)
	// param := "e5:/" + root
	// downloaddeal.OnedriveMkdir("rclone", "mkdir", param)
	// 文件处理
	filedeal.FileDeal(path, root, DomainHeader)
}
