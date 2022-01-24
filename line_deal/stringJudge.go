package linedeal

import (
	"strings"
)

// 如果字符串中包含一、二、三等
func StringJudgeHanzi(line string) bool {
	if strings.Contains(line, "一") || strings.Contains(line, "二") || strings.Contains(line, "三") || strings.Contains(line, "四") || strings.Contains(line, "五") ||
		strings.Contains(line, "六") || strings.Contains(line, "七") || strings.Contains(line, "八") || strings.Contains(line, "九") || strings.Contains(line, "十") ||
		strings.Contains(line, "十一") || strings.Contains(line, "十二") || strings.Contains(line, "十三") || strings.Contains(line, "十四") || strings.Contains(line, "十五") ||
		strings.Contains(line, "十六") || strings.Contains(line, "十七") || strings.Contains(line, "十八") || strings.Contains(line, "十九") || strings.Contains(line, "二十") ||
		strings.Contains(line, "二一") || strings.Contains(line, "二二") || strings.Contains(line, "二三") || strings.Contains(line, "二四") || strings.Contains(line, "二五") ||
		strings.Contains(line, "二六") || strings.Contains(line, "二七") || strings.Contains(line, "二八") || strings.Contains(line, "二九") || strings.Contains(line, "三十") ||
		strings.Contains(line, "三一") || strings.Contains(line, "三二") || strings.Contains(line, "三三") || strings.Contains(line, "三四") || strings.Contains(line, "三五") ||
		strings.Contains(line, "三六") || strings.Contains(line, "三七") || strings.Contains(line, "三八") || strings.Contains(line, "三九") || strings.Contains(line, "四十") ||
		strings.Contains(line, "四一") || strings.Contains(line, "四二") || strings.Contains(line, "四三") || strings.Contains(line, "四四") || strings.Contains(line, "四五") ||
		strings.Contains(line, "四六") || strings.Contains(line, "四七") || strings.Contains(line, "四八") || strings.Contains(line, "四九") || strings.Contains(line, "五十") ||
		strings.Contains(line, "五一") || strings.Contains(line, "五二") || strings.Contains(line, "五三") || strings.Contains(line, "五四") || strings.Contains(line, "五五") ||
		strings.Contains(line, "五六") || strings.Contains(line, "五七") || strings.Contains(line, "五八") || strings.Contains(line, "五九") || strings.Contains(line, "六十") ||
		strings.Contains(line, "六一") || strings.Contains(line, "六二") || strings.Contains(line, "六三") || strings.Contains(line, "六四") || strings.Contains(line, "六五") ||
		strings.Contains(line, "六六") || strings.Contains(line, "六七") || strings.Contains(line, "六八") || strings.Contains(line, "六九") || strings.Contains(line, "七十") ||
		strings.Contains(line, "七一") || strings.Contains(line, "七二") || strings.Contains(line, "七三") || strings.Contains(line, "七四") || strings.Contains(line, "七五") ||
		strings.Contains(line, "七六") || strings.Contains(line, "七七") || strings.Contains(line, "七八") || strings.Contains(line, "七九") || strings.Contains(line, "八十") ||
		strings.Contains(line, "八一") || strings.Contains(line, "八二") || strings.Contains(line, "八三") || strings.Contains(line, "八四") || strings.Contains(line, "八五") ||
		strings.Contains(line, "八六") || strings.Contains(line, "八七") || strings.Contains(line, "八八") || strings.Contains(line, "八九") || strings.Contains(line, "九十") ||
		strings.Contains(line, "九一") || strings.Contains(line, "九二") || strings.Contains(line, "九三") || strings.Contains(line, "九四") || strings.Contains(line, "九五") ||
		strings.Contains(line, "九六") || strings.Contains(line, "九七") || strings.Contains(line, "九八") || strings.Contains(line, "九九") || strings.Contains(line, "一百") {
		return true
	} else {
		return false
	}
}

// 判断字符串中不包含m3u8，即第二类文件夹
func StringJudgem3u8(line string) bool {
	if !strings.Contains(line, "m3u8") {
		return false
	} else {
		return true
	}
}
