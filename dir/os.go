package dirutil

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Ls 命令
func Ls(pattern []string) ([]string, error) {

	if len(pattern) == 0 {
		return Find([]string{"./*"})
	}

	if filepath.Base(pattern[0]) == filepath.Dir(pattern[0]) {
		return Find([]string{pattern[0] + "/*"})
	}

	return Find(pattern)
}

// Find Files……
func Find(pattern []string) ([]string, error) {

	for _, value := range pattern {
		if err := EmptyError(value); err != nil {
			return nil, err
		}
	}

	if len(pattern) == 1 {
		return filepath.Glob(pattern[0])
	} else if len(pattern) == 2 && strings.ToUpper(pattern[1]) == "R" {
		return WalkFind(pattern[0])
	}

	return nil, nil

}

// Rename ...
func Rename(file, newname string) error {

	if err := EmptyError(file, newname); err != nil {
		return err
	}

	if filepath.Base(file) == filepath.Dir(file) {
		return errors.New("要修改的文件不存在。")
	}
	return os.Rename(file, filepath.Dir(file)+string(os.PathSeparator)+newname)
}

// Time 修改
// 修改文件访问时间和修改时间
func Time(name string, atime time.Time, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}

// Mkdir 创建目录
func Mkdir(name string) error {
	return os.MkdirAll(filepath.Dir(name), os.ModePerm)
}

// Rm 删除目录及文件
func Rm(name string) error {
	return os.RemoveAll(name)
}

// ExistDir 判断路径文件夹是否存在
func ExistDir(path string) bool {
	if ok := Exist(path); ok {
		return IsDir(path)
	}
	return false
}

// Exist 判断路径文件/文件夹是否存在
func Exist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	if dir, err := os.Stat(path); err != nil {
		return dir.IsDir()
	}
	return false
}
