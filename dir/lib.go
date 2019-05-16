package dirutil

import (
	"errors"
	"os"
	"path/filepath"
)

// WalkFind Files……
// 遍历目录查找文件
func WalkFind(pattern string) ([]string, error) {

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
func EmptyError(str ...string) error {
	for _, value := range str {
		if value == "" {
			return errors.New("参数不能为空。")
		}
	}
	return nil
}
