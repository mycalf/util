package util

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Dir 初始化...
// Dir().Add("image").Add("get").Add("xxx.html").Dir
func Dir(path ...string) *utilDir {
	if len(path) == 0 {
		return &utilDir{}
	}
	return &utilDir{path[0]}
}

// Add 添加路径
func (d *utilDir) Add(path string) *utilDir {
	fmt.Println(d.dir)
	d.dir = filepath.Join(d.dir, path)
	return d
}

// Scanner ...
func (d *utilDir) Scanner() *bufio.Scanner {

	file, err := ioutil.ReadFile(d.dir)

	if err != nil {
		return nil
	}

	return bufio.NewScanner(bytes.NewReader(file))
}

// File 文件名带后嘴...
func (d *utilDir) File() string {
	return filepath.Base(d.dir)
}

// Path 文件所在目录...
func (d *utilDir) Path() string {
	return filepath.Dir(d.dir)
}

// Create 新建文件或目录
func (d *utilDir) Create() bool {

	if f, err := os.Create(d.dir); err == nil {
		defer f.Close()
		return true
	}

	return false
}

// Write 写入文件，如果不存在则创建新的文件
func (d *utilDir) Write(src []byte, add ...bool) bool {

	var f *os.File
	var err error

	defer f.Close()

	if Dir(d.Path()).Exist() == false {
		Dir(d.Path()).Mkdir()
	}

	if d.ExistFile() {
		if len(add) != 0 && add[0] == false {
			f, err = os.OpenFile(d.dir, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModeAppend)
		} else {
			f, err = os.OpenFile(d.dir, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		}
	} else {
		f, err = os.Create(d.dir)
	}

	if err == nil {
		_, err = f.Write(src)
	}

	if err == nil {
		return true
	}

	return false
}

// Filename 文件名...
func (d *utilDir) Filename(pattern ...string) string {
	switch i := len(pattern); i {
	case 1:
		return strings.Replace(d.File(), pattern[0], "", -1)
	default:
		return strings.Replace(d.File(), d.Suffix(), "", -1)
	}
}

// LastDir 文件所在目录名...
func (d *utilDir) LastDir() string {
	if d.dir == "" || d.Path() == "" {
		return ""
	}
	var dirMap []string
	if d.Suffix() == "" {
		dirMap = strings.Split(strings.Trim(d.dir, "/"), "/")
	} else {
		dirMap = strings.Split(d.Path(), "/")
	}

	return dirMap[len(dirMap)-1]

}

// Suffix 文件名
func (d *utilDir) Suffix() string {
	return filepath.Ext(d.dir)
}

// Ls 命令
func (d *utilDir) Ls(pattern ...string) ([]string, error) {
	switch i := len(pattern); i {
	case 1:
		return d.Find(pattern[0])
	case 2:
		return d.Find(pattern[0], pattern[1])
	default:
		return d.Find("/*")
	}
}

// Find Files……
func (d *utilDir) Find(pattern ...string) ([]string, error) {

	for _, value := range pattern {

		if err := emptyError(value); err != nil {
			return nil, err
		}

		if strings.ToUpper(value) == "R" {
			if len(pattern) == 2 {
				return walkFind(d.dir + "/" + pattern[0])
			}
			return walkFind(d.dir)
		}

	}

	if len(pattern) == 1 && strings.ToUpper(pattern[0]) != "R" {
		return filepath.Glob(d.dir + "/" + pattern[0])
	}

	return filepath.Glob(d.dir)

}

// Rename ...
func (d *utilDir) Rename(name string) error {

	if err := emptyError(name); err != nil {
		return err
	}

	if d.fileName(name) == d.dir {
		return errors.New("目标文件名与要修改的文件名重复。")
	}

	if d.ExistFile(d.dir) == false {
		return errors.New("要修改的文件不存在。")
	}

	if d.ExistFile(d.fileName(name)) == true {
		return errors.New("目标文件已存在。")
	}

	return os.Rename(d.dir, d.fileName(name))
}

// Retime 修改
// 修改文件访问时间和修改时间
func (d *utilDir) Retime(name string, atime time.Time, mtime time.Time) error {
	return os.Chtimes(d.fileName(name), atime, mtime)
}

// Mkdir 创建目录
func (d *utilDir) Mkdir(name ...string) error {
	if len(name) == 0 {
		return os.MkdirAll(d.dir, os.ModePerm)
	}
	return os.MkdirAll(d.fileName(name[0]), os.ModePerm)
}

// Rm 删除目录及文件
func (d *utilDir) Rm(name ...string) error {
	if len(name) == 0 {
		return os.RemoveAll(d.dir)
	}
	return os.RemoveAll(d.fileName(name[0]))
}

// ExistDir 判断路径文件夹是否存在
func (d *utilDir) ExistDir(name ...string) bool {

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

// ExistFile 判断路径文件是否存在
func (d *utilDir) ExistFile(name ...string) bool {

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
func (d *utilDir) Exist(name ...string) bool {
	var err error

	if len(name) == 0 {
		_, err = os.Stat(d.dir)
	} else {
		_, err = os.Stat(d.fileName(name[0]))
	}

	if err != nil {
		return os.IsExist(err)
	}

	return true
}

// IsDir 判断所给路径是否为文件夹
func (d *utilDir) IsDir(name ...string) bool {

	var dir os.FileInfo
	var err error

	if len(name) == 0 {
		dir, err = os.Stat(d.dir)
	} else {
		dir, err = os.Stat(d.fileName(name[0]))
	}

	if err == nil {
		return dir.IsDir()
	}

	return false
}
