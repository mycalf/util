package util

import (
	dirutil "github.com/mycalf/util/dir"
	textutil "github.com/mycalf/util/text"
)

// Text 初始化...
func Text() *textutil.Textutil {
	return textutil.Text()
}

// Dir 初始化...
// Dir().Add("image").Add("get").Add("xxx.html").String()
func Dir(path string) *dirutil.Dirutil {
	return dirutil.Dir(path)
}
