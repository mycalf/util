package util

import (
	"github.com/mycalf/util/dirutil"
	"github.com/mycalf/util/textutil"
)

// Text 初始化...
func Text(text ...string) *textutil.Textutil {
	if len(text) == 0 {
		return textutil.Text()
	}
	return textutil.Text(text[0])
}

// Dir 初始化...
func Dir(path string) *dirutil.Dirutil {
	return dirutil.Dir(path)
}
