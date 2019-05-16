package dirutil

import (
	"path/filepath"
	"strings"
)

// Dirutil 工具类型...
type Dirutil struct {
	Dir string
}

// Dir 初始化...
// Dir().Add("image").Add("get").Add("xxx.html").Dir
func Dir(path ...string) *Dirutil {
	if len(path) == 0 {
		return &Dirutil{Dir: "./"}
	}
	return &Dirutil{Dir: path[0]}
}

// Add 添加路径
func (d *Dirutil) Add(path string) *Dirutil {
	d.Dir = filepath.Join(d.Dir, path)
	return d
}

// File 文件名带后嘴...
func (d *Dirutil) File() string {
	return filepath.Base(d.Dir)
}

// Filename 文件名...
func (d *Dirutil) Filename() string {
	return strings.Replace(d.File(), d.Suffix(), "", -1)
}

// Suffix 文件名
func (d *Dirutil) Suffix() string {
	return filepath.Ext(d.Dir)
}
