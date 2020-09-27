package textutil

import (
	"bytes"
	"strings"
	"text/template"
	"unicode"
)

// Textutil 工具类型...
type Textutil struct {
	Text string
}

// Text 初始化...
func Text(str ...string) *Textutil {
	if len(str) == 0 {
		return &Textutil{}
	}
	return &Textutil{Text: str[0]}
}

// Initials 英文首字母大写 ...
func (t *Textutil) Initials() string {
	for i, v := range t.Text {
		return string(unicode.ToUpper(v)) + t.Text[i+1:]
	}
	return t.Text
}

// Lower 字符串全部小写 ...
func (t *Textutil) Lower() string {
	return strings.ToLower(t.Text)
}

// Upper 字符串全部大写 ...
func (t *Textutil) Upper() string {
	return strings.ToUpper(t.Text)
}

// UpperID 查找ID在字符串是否出现在字符串的最后，如果出现，则将整个字符串改为大写 ...
func (t *Textutil) UpperID() string {

	t.Text = t.Initials()

	if len(t.Text)-len("id") == strings.Index(t.Lower(), "id") {
		if len(t.Text) <= 6 {
			return t.Upper()
		}
		return strings.Replace(t.Text, "id", "ID", -1)
	}

	return t.Text
}

// Trim 去除开始及结束出现的字符 ...
func (t *Textutil) Trim(sep string) string {
	return strings.Trim(t.Text, sep)
}

// Split 根据字符串进行文本分割
func (t *Textutil) Split(sep string) []string {
	return strings.Split(t.Text, sep)
}

// SplitPlace 根据字符串的位置进行分割
// Text("abcdefg").SpltPlace([]int{1,3,4})
// Out: []string{"a", "bc", "d", "efg"}
func (t *Textutil) SplitPlace(sep []int) []string {
	var a []string
	b := Text()
	for k, v := range []rune(t.Text) {
		b.Add(string(v))
		for _, i := range sep {
			if i == k+1 {
				a = append(a, b.Text)
				b = Text()
			}
		}

		if len(t.Text) == k+1 {
			a = append(a, b.Text)
		}
	}
	return a
}

// Add 在文本中追加文字
func (t *Textutil) Add(text string) *Textutil {
	bufferString := bytes.NewBufferString(t.Text)
	bufferString.WriteString(text)
	t.Text = bufferString.String()
	return t
}

// Parse 根据模版解析文本...
// t.Text : Hell{{.d}} Wor{{.o}}
// data : map[string]string{ "d": "o", "o":"d" }
// t.Text : Hello {{.}}
// data : Word
// t.Text : {{.H}} {{.W}}
// &data{H: "Hello", W: "Word"}
// -> Hello Word
func (t *Textutil) Parse(data interface{}) string {
	var text bytes.Buffer
	if err := template.Must(template.New("").Parse(t.Text)).Execute(&text, data); err == nil {
		return text.String()
	}
	return ""
}

// ChineseNumber 英文数字转为中文数字
func (t *Textutil) ChineseNumber(mode bool) string {
	a := strings.Split(t.Text, ".")

	if len(a) == 1 {
		return t.chineseInt(mode)
	}

	if len(a) == 2 {
		b := Text()
		b.Add(Text(a[0]).chineseInt(mode))
		b.Add(chineseDot(mode))
		b.Add(Text(a[1]).chineseFloat(mode))
		return b.Text
	}
	return ""
}
