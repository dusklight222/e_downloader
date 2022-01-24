package linedeal

import (
	"fmt"
)

// 进行字符串的筛选
func StringDeal(localRoot string, root string, DomainHeader string, line string, localName string, localUrl string) (string, string, string) {
	// 继承之前的
	// 当前路径
	var nowRoot string = localRoot
	// 当前文件名
	var nowName string = localName
	// 当前url
	var nowUrl string = localUrl
	// 判断字符串中是否包含一、二、三等等
	if StringJudgeHanzi(line) {
		// 创建对应的文件夹
		nowRoot = StringDealHanzi(root, line)
		fmt.Println(nowRoot)
	} else if !StringJudgem3u8(line) {
		// 说明当前的字符串是视频的名称
		// 将文件的绝对路径名称返回
		nowName = StringDealName(localRoot, line)
		fmt.Println(nowName)
	} else {
		// 对包含m3u8的字符串进行处理
		nowUrl = StringDealM3u8(DomainHeader, line)
		fmt.Println(nowUrl)
		// os.Exit(-1)
	}
	return nowRoot, nowName, nowUrl
}
