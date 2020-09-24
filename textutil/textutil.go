package textutil

import (
	"bytes"
	"strings"
	"text/template"
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
func (t *Textutil) Initials() string  {
	return Strings.ToTitle(t.Text)
}

// Lower 字符串全部小写 ...
func (t *Textutil) Lower() string  {
	return Strings.ToLower(t.Text)
}
// Upper 字符串全部大写 ...
func (t *Textutil) Upper() string  {
	return Strings.ToUpper(t.Text)
}



// Split 根据字符串进行文本分割
func (t *Textutil) Split(sep string) []string {
	return strings.Split(t.Text, sep)
}

// SplitPlace 根据位置进行文本分割
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
