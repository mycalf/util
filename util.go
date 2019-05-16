package util

import (
	dirutil "github.com/mycalf/util/dir"
	"github.com/mycalf/util/text"
)

// Text 初始化...
func Text() *text.Textutil {
	return &text.Textutil{}
}

// Dir 初始化...
// Dir().Add("image").Add("get").Add("xxx.html").String()
func Dir() *dirutil.Dirutil {
	return &dirutil.Dirutil{}
}
