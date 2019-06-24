package dirutil

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
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

// Create 新建文件
func (d *Dirutil) Create() bool {

	if f, err := os.Create(d.Dir); err == nil {
		defer f.Close()
		return true
	}

	return false
}

// Filename 文件名...
func (d *Dirutil) Filename() string {
	return strings.Replace(d.File(), d.Suffix(), "", -1)
}

// Suffix 文件名
func (d *Dirutil) Suffix() string {
	return filepath.Ext(d.Dir)
}

// Ls 命令
func (d *Dirutil) Ls(pattern ...string) ([]string, error) {

	if len(pattern) == 1 {
		return d.Find(d.Dir + pattern[0])
	}

	if filepath.Base(d.Dir) == filepath.Dir(d.Dir) {
		return d.Find("/*")
	}

	return d.Find()
}

// Find Files……
func (d *Dirutil) Find(pattern ...string) ([]string, error) {

	for _, value := range pattern {

		if err := emptyError(value); err != nil {
			return nil, err
		}

		if strings.ToUpper(value) == "R" {
			if len(pattern) == 2 {
				return walkFind(d.Dir + "/" + pattern[0])
			}
			return walkFind(d.Dir)
		}

	}

	if len(pattern) == 1 && strings.ToUpper(pattern[0]) != "R" {
		return filepath.Glob(d.Dir + "/" + pattern[0])
	}

	return filepath.Glob(d.Dir)

}

// Rename ...
func (d *Dirutil) Rename(name string) error {

	if err := emptyError(name); err != nil {
		return err
	}

	if d.fileName(name) == d.Dir {
		return errors.New("目标文件名与要修改的文件名重复。")
	}

	if d.ExistFile(d.Dir) == false {
		return errors.New("要修改的文件不存在。")
	}

	if d.ExistFile(d.fileName(name)) == true {
		return errors.New("目标文件已存在。")
	}

	return os.Rename(d.Dir, d.fileName(name))
}

// Time 修改
// 修改文件访问时间和修改时间
func (d *Dirutil) Time(name string, atime time.Time, mtime time.Time) error {
	return os.Chtimes(d.fileName(name), atime, mtime)
}

// Mkdir 创建目录
func (d *Dirutil) Mkdir(name ...string) error {
	if len(name) == 0 {
		return os.MkdirAll(d.Dir, os.ModePerm)
	}
	return os.MkdirAll(d.fileName(name[0]), os.ModePerm)
}

// Rm 删除目录及文件
func (d *Dirutil) Rm(name ...string) error {
	if len(name) == 0 {
		return os.RemoveAll(d.Dir)
	}
	return os.RemoveAll(d.fileName(name[0]))
}

// ExistDir 判断路径文件夹是否存在
func (d *Dirutil) ExistDir(name ...string) bool {

	if len(name) == 0 {
		if ok := d.Exist(); ok {
			return d.IsDir()
		}
	}

	if ok := d.Exist(name[0]); ok {
		return d.IsDir(name[0])
	}

	return false
}

// ExistFile 判断路径文件夹是否存在
func (d *Dirutil) ExistFile(name ...string) bool {

	if len(name) == 0 {
		if ok := d.Exist(); ok {
			return d.IsDir() == false
		}
	} else {
		if ok := d.Exist(name[0]); ok {
			return d.IsDir(name[0]) == false
		}
	}

	return false
}

// Exist 判断路径文件/文件夹是否存在
func (d *Dirutil) Exist(name ...string) bool {
	var err error

	if len(name) == 0 {
		_, err = os.Stat(d.Dir)
	} else {
		_, err = os.Stat(d.fileName(name[0]))
	}

	if err != nil && os.IsExist(err) {
		return true
	} else if err != nil {
		return false
	}

	return true
}

// IsDir 判断所给路径是否为文件夹
func (d *Dirutil) IsDir(name ...string) bool {

	var dir os.FileInfo
	var err error

	if len(name) == 0 {
		dir, err = os.Stat(d.Dir)
	} else {
		dir, err = os.Stat(d.fileName(name[0]))
	}

	if err == nil {
		return dir.IsDir()
	}

	return false
}
