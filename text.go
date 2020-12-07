package util

import (
	"bytes"
	"html/template"
	"strings"
	"unicode"
)

// Text 初始化...
func Text(str ...string) *UText {
	if len(str) == 0 {
		return &UText{}
	}
	return &UText{text: str[0]}
}

// Get 获取文本...
func (t *UText) Get() string {
	return t.text
}

// Find 搜索文本...
func (t *UText) Find(dst string) bool {
	return strings.Contains(t.text, dst)
}

// Replace 搜索文本...
func (t *UText) Replace(src, dst string) string {
	return strings.Replace(t.text, src, dst, -1)
}

// Space 加入空格 ...
func (t *UText) Space() *UText {
	return t.Add(" ")
}

// Enter 加入回车 ...
func (t *UText) Enter() *UText {
	return t.Add("\n")
}

// Initials 英文首字母大写 ...
func (t *UText) Initials() string {
	return Initials(t.text)
}

// Lower 字符串全部小写 ...
func (t *UText) Lower() string {
	return Lower(t.text)
}

// Upper 字符串全部大写 ...
func (t *UText) Upper() string {
	return Upper(t.text)
}

// UpperID 查找ID在字符串是否出现在字符串的最后，如果出现，则将整个字符串改为大写 ...
func (t *UText) UpperID() string {
	return UpperID(t.text)
}

// Trim 去除开始及结束出现的字符 ...
func (t *UText) Trim(sep string) string {
	return Trim(t.text, sep)
}

// Split 根据字符串进行文本分割
func (t *UText) Split(sep string) []string {
	return strings.Split(t.text, sep)
}

// SplitPlace 根据字符串的位置进行分割
// Text("abcdefg").SpltPlace([]int{1,3,4})
// Out: []string{"a", "bc", "d", "efg"}
func (t *UText) SplitPlace(sep []int) []string {
	var a []string
	b := Text()
	for k, v := range []rune(t.text) {
		b.Add(string(v))
		for _, i := range sep {
			if i == k+1 {
				a = append(a, b.text)
				b = Text()
			}
		}

		if len(t.text) == k+1 {
			a = append(a, b.text)
		}
	}
	return a
}

// Add 在文本中追加文字
func (t *UText) Add(text string) *UText {
	bufferString := bytes.NewBufferString(t.text)
	bufferString.WriteString(text)
	t.text = bufferString.String()
	return t
}

// Parse 根据模版解析文本...
// t.text : Hell{{.d}} Wor{{.o}}
// data : map[string]string{ "d": "o", "o":"d" }
// t.text : Hello {{.}}
// data : Word
// t.text : {{.H}} {{.W}}
// &data{H: "Hello", W: "Word"}
// -> Hello Word
func (t *UText) Parse(data interface{}) string {
	var text bytes.Buffer
	if err := template.Must(template.New("").Parse(t.text)).Execute(&text, data); err == nil {
		return text.String()
	}
	return ""
}

// ChineseNumber 英文数字转为中文数字
func (t *UText) ChineseNumber(mode bool) string {
	a := t.Split(".")

	if len(a) == 1 {
		return t.chineseInt(mode)
	}

	if len(a) == 2 {
		b := Text()
		b.Add(Text(a[0]).chineseInt(mode))
		b.Add(chineseDot(mode))
		b.Add(Text(a[1]).chineseFloat(mode))
		return b.text
	}
	return ""
}

// Initials 英文首字母大写 ...
func Initials(src string) string {
	for i, v := range src {
		return string(unicode.ToUpper(v)) + src[i+1:]
	}
	return src
}

// Lower 字符串全部小写 ...
func Lower(src string) string {
	return strings.ToLower(src)
}

// Upper 字符串全部大写 ...
func Upper(src string) string {
	return strings.ToUpper(src)
}

// UpperID 查找ID在字符串是否出现在字符串的最后，如果出现，则将整个字符串改为大写 ...
func UpperID(src string) string {

	src = Initials(src)

	if len(src)-len("id") == strings.Index(Lower(src), "id") {
		if len(src) <= 6 {
			return Upper(src)
		}
		return strings.Replace(src, "id", "ID", -1)
	}

	return src
}

// Trim 去除开始及结束出现的字符 ...
func Trim(src, sep string) string {
	return strings.Trim(src, sep)
}
