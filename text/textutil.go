package textutil

import (
	"bytes"
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
