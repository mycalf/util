package util

import (
	"time"

	dirutil "github.com/mycalf/util/dir"
)

// Ls 命令
func Ls(pattern ...string) ([]string, error) {
	return dirutil.Ls(pattern)
}

// Find Files……
func Find(pattern ...string) ([]string, error) {
	return dirutil.Find(pattern)

}

// Rename ...
func Rename(file, newname string) error {
	return dirutil.Rename(file, newname)
}

// Time 修改
// 修改文件访问时间和修改时间
func Time(name string, atime time.Time, mtime time.Time) error {
	return dirutil.Time(name, atime, mtime)
}

// Mkdir 创建目录
func Mkdir(path string) error {
	return dirutil.Mkdir(path)
}

// Rm 删除目录及文件
func Rm(path string) error {
	return dirutil.Rm(path)
}

// ExistDir 判断路径文件夹是否存在
func ExistDir(path string) bool {
	return dirutil.ExistDir(path)
}

// Exist 判断路径文件/文件夹是否存在
func Exist(path string) bool {
	return dirutil.Exist(path)
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	return dirutil.IsDir(path)
}
