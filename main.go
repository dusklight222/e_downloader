package main

import (
	filedeal "e_downloader/file_deal"
	"fmt"
	"os"
)

var (
	// 用于更换的域名头
	DomainHeader string = "1252524126.vod2.myqcloud.com"
)

func main() {
	list := os.Args
	if len(list) != 3 {
		fmt.Println("请输入您的URL文件 下载器")
		return
	}
	path := list[1]
	downloaderPath := list[2]
	// 创建根目录文件夹
	root, temp := filedeal.RootMkdir(path)
	if temp {
		// 文件处理
		filedeal.FileDeal(path, root, DomainHeader, downloaderPath)
	} else {
		return
	}

}
