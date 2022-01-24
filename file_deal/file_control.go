package filedeal

import (
	"bufio"
	"e_downloader_client/download_deal"
	"e_downloader_client/line_deal"
	"fmt"
	"io"
	"os"
	"strings"
)

// 文件操作
// 文件路径 根路径 当前路径 域名头
func FileDeal(path string, root string, domainHeader string) {
	// 当前文件名
	var localName string
	// 当前URL
	var localUrl string
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
		// fmt.Println(line)
		// nowRoot, nowName, nowUrl
		localRoot, localName, localUrl = linedeal.StringDeal(localRoot, root, domainHeader, line, localName, localUrl)
		// 下载视频
		if localUrl != "" {
			if !downloaddeal.UrlDownload(localUrl, localName, localRoot) {
				fmt.Println("失败文件名为：", localName)
				fmt.Println("失败URL为", localUrl)
				os.Exit(-1)
				// continue
			}
		}

	}
}
