package text

import (
	"bytes"
	"text/template"
)

// Parse 根据模版解析文本...
// letter : Hell{{.d}} Wor{{.o}}
// data : map[string]string{ "d": "o", "o":"d" }
// letter : Hello {{.}}
// data : Word
// letter : {{.H}} {{.W}}
// &data{H: "Hello", W: "Word"}
// -> Hello Word
func Parse(letter string, data interface{}) string {
	var text bytes.Buffer
	err := template.Must(template.New("").Parse(letter)).Execute(&text, data)
	if err == nil {
		return text.String()
	}
	return ""
}
