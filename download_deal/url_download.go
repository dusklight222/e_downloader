package downloaddeal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func UrlDownload(url string, localName string, localRoot string, downloaderPath string) bool {
	// 下载url
	command := downloaderPath
	param1 := "-u=" + url
	// 下载完成后的main.ts会生成在当前工作目录下
	root, _ := os.Getwd()
	param2 := "-o=" + root
	// M3u8DownloaderUse()
	if !M3u8DownloaderUse(command, param1, param2) {
		// fmt.Println("文件下载失败")
		return false
	}
	// 对文件进行改名
	// 获取当前工作路径
	path := "./main.ts"
	localName = localName + ".mp4"
	if !FileNameChange(path, localName) {
		// fmt.Println("文件改名失败")
		return false
	}
	return true
}

// 调用m3u8下载器下载url
func M3u8DownloaderUse(command string, param1 string, param2 string) bool {
	shell_command := exec.Command(command, param1, param2)
	// fmt.Println(shell_command.Args)
	stdout, err := shell_command.StdoutPipe()
	if err != nil {
		// fmt.Println(err)
		// log.Println(err)
		return false
		// os.Exit(-1)
	}
	shell_command.Start()
	//创建流来接收读取内容
	reader := bufio.NewReader(stdout)
	for {
		//按行读取
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(line)
	}
	shell_command.Wait()
	return true
}

// 修改文件名称
func FileNameChange(path string, dstpath string) bool {
	err := os.Rename(path, dstpath)
	if err != nil {
		// fmt.Println("文件名称更改失败")
		// os.Exit(-1)
		return false
	} else {
		return true
	}
}
