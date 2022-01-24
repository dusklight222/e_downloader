package filedeal

import (
	"bufio"
	"e_downloader/download_deal"
	"e_downloader/line_deal"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 创建根目录文件夹
func RootMkdir(path string) (string, bool) {
	// 获取文件名
	num1 := strings.LastIndex(path, ".txt")
	// 获取文件名，以此作为根目录
	urlName := path[:num1]
	// 获取当前工作目录
	rootTemp, _ := os.Getwd()
	// 根目录进行拼接
	var rootBuild strings.Builder
	rootBuild.WriteString(rootTemp)
	rootBuild.WriteRune(filepath.Separator)
	rootBuild.WriteString(urlName)
	root := rootBuild.String()
	// 创建根目录文件名
	err := os.Mkdir(root, os.ModePerm)
	if err != nil {
		fmt.Println("文件夹创建失败，请检查权限问题")
		return path, false
	} else {
		// 成功创建，返回根目录位置
		return root, true
	}
}

// 文件操作
// 文件路径 根路径 当前路径 域名头
func FileDeal(path string, root string, domainHeader string, downloaderPath string) {
	// 当前文件名
	var localName string
	// 当前URL
	var localUrl string
	// 临时URL，防止多次下载同一链接
	var tempUrl string
	// 当前路径
	var localRoot string = root
	filePath, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开失败")
		os.Exit(-1)
	}
	// 延迟关闭文件
	defer filePath.Close()
	// 读取文件内容
	reader1 := bufio.NewReader(filePath)
	for {
		line, err := reader1.ReadString('\n')
		if io.EOF == nil {
			fmt.Println("文件读取结束")
			os.Exit(-1)
		}
		if err != nil {
			fmt.Println("文件内容读取错误")
			os.Exit(-1)
		}
		// 去掉其中的换行符
		line = strings.Replace(line, "\r\n", "", -1)
		// 打印文件内容
		// nowRoot, nowName, nowUrl
		localRoot, localName, localUrl = linedeal.StringDeal(localRoot, root, domainHeader, line, localName, localUrl)
		// 下载视频
		if localUrl != tempUrl {
			if !downloaddeal.UrlDownload(localUrl, localName, localRoot, downloaderPath) {
				fmt.Println("失败文件名为：", localName)
				fmt.Println("失败URL为", localUrl)
				continue
			}
		}
		// 将已经下载过的URL赋值给临时变量，防止多次下载同一链接
		tempUrl = localUrl
	}
}
