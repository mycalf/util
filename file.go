package util

import (
	"os"
	"path/filepath"
	"strings"
)

// File 文件全名,带文件类型的文件名 ...
// In: app/conf/users.db.yaml
// out: users.db.yaml
func (d *utilOS) File() string {
	return d.file
}

// FileName 不带文件类型的文件名 ...
// In: app/conf/users.db.yaml
// out: users.db
func (d *utilOS) FileName(pattern ...string) string {
	switch i := len(pattern); i {
	case 1:
		return strings.Replace(d.File(), pattern[0], "", -1)
	default:
		return strings.Replace(d.File(), d.Suffix(), "", -1)
	}
}

// Suffix 文件后缀 ...
func (d *utilOS) Suffix() string {
	return filepath.Ext(d.Path())
}

// Create 新建文件或目录
func (d *utilOS) Create() bool {
	if f, err := os.Create(d.Path()); err == nil {
		defer f.Close()
		return true
	}
	return false
}

// IsFile 判断是否是文件 ...
func (d *utilOS) IsFile() bool {
	return d.Exist() && d.IsDir() == false

}

// Write 写入文件，如果不存在则创建新的文件
func (d *utilOS) Write(src []byte, add ...bool) bool {

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
// func (d *utilOS) LoadYAML(in ...string) (*viper.Viper, bool) {
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
