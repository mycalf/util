package util

import (
	"os"
	"path/filepath"
	"strings"
)

// OS 初始化...
// OS().Add("image").Add("get").Add("xxx.html").Dir
func OS(path ...string) *utilOS {
	d := &utilOS{}
	if len(path) > 0 {
		d.path = filepath.Join(d.path, path[0])
		d.dir, d.file = filepath.Split(d.path)
	}
	return d
}

// Add 添加路径 ...
func (d *utilOS) Add(path string) *utilOS {
	d.path = filepath.Join(d.path, path)
	d.dir, d.file = filepath.Split(d.path)
	return d
}

// Path 当前路径 ...
func (d *utilOS) Path() string {
	return d.path
}

// Dir 当前目录 ...
func (d *utilOS) Dir() string {
	return d.dir
}

// IsDir 判断所给路径是否为文件夹
func (d *utilOS) IsDir(path ...string) bool {
	var dir os.FileInfo
	var err error
	p := d.Path()
	dir, err = os.Stat(p)
	if err == nil {
		return dir.IsDir()
	}
	return false
}

// InDir 文件所在目录，不包含路径...
func (d *utilOS) InDir() string {
	if d.Dir() == "" || d.Path() == "" {
		return ""
	}
	out := strings.Split(d.Dir(), "/")
	return out[len(out)-1]

}

// MkDir 创建目录
func (d *utilOS) MkDir() error {
	return os.MkdirAll(d.Dir(), os.ModePerm)
}

// Exist 判断路径文件/文件夹是否存在
func (d *utilOS) Exist(path ...string) bool {
	p := d.Path()
	if len(path) > 0 {
		p = path[0]
	}
	if _, err := os.Stat(p); err != nil {
		return os.IsExist(err)
	}
	return true
}

// Exist 判断路径文件/文件夹是否存在
func (d *utilOS) DirExist() bool {
	return d.Exist(d.Dir())
}

// Ls 命令
func (d *utilOS) Ls(pattern ...string) ([]string, error) {
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
func (d *utilOS) Find(pattern ...string) ([]string, error) {

	for _, value := range pattern {

		if err := emptyError(value); err != nil {
			return nil, err
		}

		if strings.ToUpper(value) == "R" {
			if len(pattern) == 2 {
				return walkFind(d.Path() + string(os.PathSeparator) + pattern[0])
			}
			return walkFind(d.Path())
		}

	}

	if len(pattern) == 1 && strings.ToUpper(pattern[0]) != "R" {
		return filepath.Glob(d.Path() + string(os.PathSeparator) + pattern[0])
	}

	return filepath.Glob(d.Path())

}

// --------

// // Rename ...
// func (d *utilOS) Rename(name string) error {

// 	if err := emptyError(name); err != nil {
// 		return err
// 	}

// 	if d.fileName(name) == d.Path() {
// 		return errors.New("目标文件名与要修改的文件名重复。")
// 	}

// 	if d.IsFile(d.Path()) == false {
// 		return errors.New("要修改的文件不存在。")
// 	}

// 	if d.IsFile(d.fileName(name)) == true {
// 		return errors.New("目标文件已存在。")
// 	}

// 	return os.Rename(d.Path(), d.fileName(name))
// }

// // Retime 修改
// // 修改文件访问时间和修改时间
// func (d *utilOS) Retime(name string, a time.Time, m time.Time) error {
// 	return os.Chtimes(d.fileName(name), a, m)
// }

// // Rm 删除目录及文件
// func (d *utilOS) Rm(name ...string) error {
// 	if len(name) == 0 {
// 		return os.RemoveAll(d.Path())
// 	}
// 	return os.RemoveAll(d.fileName(name[0]))
// }

// // Scanner ...
// func (d *utilOS) Scanner() *bufio.Scanner {

// 	file, err := ioutil.ReadFile(d.Path())

// 	if err != nil {
// 		return nil
// 	}

// 	return bufio.NewScanner(bytes.NewReader(file))
// }
