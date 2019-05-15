package util

import (
	"github.com/mycalf/util/text"
)

// IsHan 判断是否为中文...
func IsHan(str string) bool {
	return text.IsHan(str)
}

// Parse 根据模版解析文本...
// letter : Hell{{.d}} Wor{{.o}}
// data : map[string]string{ "d": "o", "o":"d" }
// letter : Hello {{.}}
// data : Word
// letter : {{.H}} {{.W}}
// &data{H: "Hello", W: "Word"}
// -> Hello Word
func Parse(letter string, data interface{}) string {
	return text.Parse(letter, data)
}

// NumberToChinese ...
// 数字转汉字
// 第二个参数为大小写开关
func NumberToChinese(str string) string {
	return text.NumberToChinese(str)
}
