package downloaddeal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func UrlDownload(url string, localName string, localRoot string) bool {
	// 下载url
	command := "m3u8-download.exe"
	// var params []string
	param1 := "-u=" + url
	root, _ := os.Getwd()
	param2 := "-o=" + root
	// M3u8DownloaderUse()
	if !M3u8DownloaderUse(command, param1, param2) {
		// fmt.Println("文件下载失败")
		return false
	}
	// 对文件进行改名
	// 获取当前工作路径
	path := "main.ts"
	localName = localName + ".mp4"
	if !FileNameChange(path, localName) {
		// fmt.Println("文件改名失败")
		return false
	}
	// 将文件上传至onedrive
	var param3Builder strings.Builder
	param3Builder.WriteString(root)
	param3Builder.WriteRune(filepath.Separator)
	param3Builder.WriteString(localName)
	param3 := param3Builder.String()
	var param4Builder strings.Builder
	param4Builder.WriteString("e5:/")
	param4Builder.WriteString(localRoot)
	param4Builder.WriteRune(filepath.Separator)
	// param4Builder.WriteString(localName)
	param4 := param4Builder.String()
	param5 := "-P"
	if !OnedriveUploader("rclone", "copy", param3, param4, param5) {
		// 文件上传失败
		return false
	}
	// 上传完成后，删除文件
	os.Remove(localName)
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

// param1:rclone
// param2:mkdir
// param3:path
// 在onedrive上创建文件夹
func OnedriveMkdir(command string, param1 string, param2 string) bool {
	shell_command := exec.Command(command, param1, param2)
	// fmt.Println(shell_command.Args)
	stdout, err := shell_command.StdoutPipe()
	if err != nil {
		return false
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

// 参数1：rclone
// 参数2：copy
// 参数3：本地路径
// 参数4：目标路径
// 参数5:-P
// 将视频上传至onedrive
func OnedriveUploader(command string, param1 string, param2 string, param3 string, param4 string) bool {
	shell_command := exec.Command(command, param1, param2, param3, param4)
	// fmt.Println(shell_command.Args)
	stdout, err := shell_command.StdoutPipe()
	if err != nil {
		return false
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
