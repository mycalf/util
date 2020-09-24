package dirutil

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
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

// Scanner ...
func (d *Dirutil) Scanner() *bufio.Scanner {

	file, err := ioutil.ReadFile(d.Dir)

	if err != nil {
		return nil
	}

	return bufio.NewScanner(bytes.NewReader(file))
}

// File 文件名带后嘴...
func (d *Dirutil) File() string {
	return filepath.Base(d.Dir)
}

// Path 文件所在目录...
func (d *Dirutil) Path() string {
	return filepath.Dir(d.Dir)
}

// Create 新建文件或目录
func (d *Dirutil) Create() bool {

	if f, err := os.Create(d.Dir); err == nil {
		defer f.Close()
		return true
	}

	return false
}

// Write 写入文件，如果不存在则创建新的文件
func (d *Dirutil) Write(src []byte, add ...bool) bool {

	var f *os.File
	var err error

	defer f.Close()

	if Dir(d.Path()).Exist() == false {
		Dir(d.Path()).Mkdir()
	}

	if d.ExistFile() {
		if len(add) != 0 && add[0] == false {
			f, err = os.OpenFile(d.Dir, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModeAppend)
		} else {
			f, err = os.OpenFile(d.Dir, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		}
	} else {
		f, err = os.Create(d.Dir)
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
func (d *Dirutil) Filename(pattern ...string) string {
	switch i := len(pattern); i {
	case 1:
		return strings.Replace(d.File(), pattern[0], "", -1)
	default:
		return strings.Replace(d.File(), d.Suffix(), "", -1)
	}
}

// LastDir 文件所在目录名...
func (d *Dirutil) LastDir() string {
	if d.Dir == "" || d.Path() == "" {
		return ""
	}
	var dirMap []string
	if d.Suffix() == "" {
		dirMap = strings.Split(strings.Trim(d.Dir, "/"), "/")
	} else {
		dirMap = strings.Split(d.Path(), "/")
	}

	return dirMap[len(dirMap)-1]

}

// Suffix 文件名
func (d *Dirutil) Suffix() string {
	return filepath.Ext(d.Dir)
}

// Ls 命令
func (d *Dirutil) Ls(pattern ...string) ([]string, error) {
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

// Retime 修改
// 修改文件访问时间和修改时间
func (d *Dirutil) Retime(name string, atime time.Time, mtime time.Time) error {
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

// ExistFile 判断路径文件是否存在
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

	if err != nil {
		return os.IsExist(err)
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

// WalkFind Files……
// 遍历目录查找文件
func walkFind(pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(filepath.Dir(pattern), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {

			globPattern := path + string(os.PathSeparator) + filepath.Base(pattern)

			if fileArray, err := filepath.Glob(globPattern); len(fileArray) != 0 && err == nil {
				for _, file := range fileArray {
					matches = append(matches, file)
				}
			} else {
				return err
			}
		}
		return nil
	})
	return matches, err
}

// EmptyError 错误信息
func emptyError(str ...string) error {
	for _, value := range str {
		if value == "" {
			return errors.New("参数不能为空。")
		}
	}
	return nil
}

// fileName 设置文件名
func (d *Dirutil) fileName(name string) string {
	return filepath.Dir(d.Dir) + string(os.PathSeparator) + name
}
