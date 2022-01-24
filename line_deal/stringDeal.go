package linedeal

import (
	"path/filepath"
	"strings"
)

// 处理包含汉字的字符串
func StringDealHanzi(root string, line string) string {
	// 在当前目录创建一个文件夹，名称为传入的值
	// 返回当前的工作目录
	var stringBuild strings.Builder
	stringBuild.WriteString(root)
	stringBuild.WriteRune(filepath.Separator)
	stringBuild.WriteString(line)
	workdir := stringBuild.String()
	// err := os.Mkdir(workdir, os.ModePerm)
	// if err != nil {
	// 	// fmt.Println("目录创建失败")
	// 	os.Exit(-1)
	// }
	//在onedrive上创建目录
	// command := "rclone"
	// param1 := "mkdir"
	// param2 := "e5:/" + line
	// if !downloaddeal.OnedriveMkdir(command, param1, param2) {
	// 	os.Exit(-1)
	// }
	return workdir
}

// 处理包含m3u8的字符串
func StringDealM3u8(DomainHeader string, line string) string {
	// 首先将字符串进行截取，去掉?后面的
	// https://encrypt-k-vod.xet.tech/529d8d60vodtransbj1252524126/ed4d88873701925921137477743/drm/v.f421218.m3u8?t=61eabf8b&exper=0&us=vLKxWlQ03UJS&whref=www.1cxy.net&sign=ffb5d2431b8cc9cf56bbb3b85ec46728
	// 用于记录?的最后位置
	num1 := strings.LastIndex(line, "?")
	urlTemp := line[:num1]
	// 记录域名头的位置
	num2 := strings.LastIndex(urlTemp, "encrypt")
	num3 := strings.LastIndex(urlTemp, "tech")
	var urlBuild strings.Builder
	urlBuild.WriteString(urlTemp[:num2])
	urlBuild.WriteString(DomainHeader)
	urlBuild.WriteString(urlTemp[num3+4:])
	return urlBuild.String()
}

// 根据当前路径和文件名生成下载文件的绝对路径
func StringDealName(localRoot string, localName string) string {
	// var pathBuilder strings.Builder
	// pathBuilder.WriteString(localRoot)
	// pathBuilder.WriteRune(filepath.Separator)
	// pathBuilder.WriteString(localName)
	// pathBuilder.WriteString(".mp4")
	// path := pathBuilder.String()
	// param := "e5:/" + localName
	// downloaddeal.OnedriveMkdir("rclone", "mkdir", param)
	// return path
	return localName
}
