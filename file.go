package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// File 文件全名,带文件类型的文件名 ...
// In: app/conf/users.db.yaml
// out: users.db.yaml
func (d *UOS) File() string {
	return d.file
}

// FileName 不带文件类型的文件名 ...
// In: app/conf/users.db.yaml
// out: users.db
func (d *UOS) FileName(pattern ...string) string {
	switch i := len(pattern); i {
	case 1:
		return strings.Replace(d.File(), pattern[0], "", -1)
	default:
		return strings.Replace(d.File(), d.Suffix(), "", -1)
	}
}

// Suffix 文件后缀 ...
func (d *UOS) Suffix() string {
	return filepath.Ext(d.Path())
}

// Create 新建文件或目录
func (d *UOS) Create() bool {
	if f, err := os.Create(d.Path()); err == nil {
		defer f.Close()
		return true
	}
	return false
}

// IsFile 判断是否是文件 ...
func (d *UOS) IsFile() bool {
	return d.Exist() && d.IsDir() == false

}

// Cat 查看文件.
func (d *UOS) Cat() string {
	if src, err := ioutil.ReadFile(d.path); err == nil {
		return string(src)
	}
	return ""
}

// Read File 非UTF8格式自动转换为UTF8.
func (d *UOS) Read() (string, bool) {
	return converter(d.Cat())
}

// Write 写入文件，如果不存在则创建新的文件
func (d *UOS) Write(src []byte, add ...bool) bool {

	var f *os.File
	var err error

	defer f.Close()

	if d.DirExist() == false {
		d.MkDir()
	}

	if d.IsFile() {
		if len(add) != 0 && add[0] == false {
			f, err = os.OpenFile(d.Path(), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModeAppend)
		} else {
			f, err = os.OpenFile(d.Path(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		}
	} else {
		f, err = os.Create(d.Path())
	}

	if err == nil {
		_, err = f.Write(src)
	}

	if err == nil {
		return true
	}

	return false
}

// // LoadYAML : Load yyt YAML
// func (d *UOS) LoadYAML(in ...string) (*viper.Viper, bool) {
// 	in = append(in, d.Path())
// 	process, err := files.NewSortedFilesFromPaths(in, files.SymlinkAllowOpts{})

// 	if err != nil {
// 		return nil, false
// 	}

// 	opts := template.NewOptions()

// 	yyt := opts.RunWithFiles(template.TemplateInput{Files: process}, core.NewPlainUI(false))

// 	if yyt.Err != nil || len(yyt.DocSet.Items) < 1 {
// 		return nil, false
// 	}

// 	yamlBytes, err := yyt.DocSet.Items[0].AsYAMLBytes()

// 	if err != nil {
// 		return nil, false
// 	}

// 	out := viper.New()
// 	out.SetConfigType("yaml")
// 	out.ReadConfig(bytes.NewBuffer(yamlBytes))

// 	return out, true
// }
