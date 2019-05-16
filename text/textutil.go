package textutil

import (
	"bytes"
)

// Textutil 工具类型...
type Textutil struct {
	Text string
}

// Text 初始化...
func Text() *Textutil {
	return &Textutil{}
}

// Add 在文本中追加文字
func (t *Textutil) Add(text string) *Textutil {
	bufferString := bytes.NewBufferString(t.Text)
	bufferString.WriteString(text)
	t.Text = bufferString.String()
	return t
}
