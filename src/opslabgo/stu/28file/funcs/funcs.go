package funcs

import (
	"io/ioutil"
	"os"
	"strings"
	"path/filepath"
)

/**
 获取指定目录下的所有文件，不进入下一级目录搜索,可以后缀过滤
 */
func List_dir(dir_path string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dir_path)

	if err != nil {
		return nil, err
	}

	pthsep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)

	for _, fi := range dir {
		if fi.IsDir() {
			//忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dir_path + pthsep + fi.Name())
		}
	}
	return files, nil
}

func Walk_dir(dir_path string, suffix string) (files[]string, err error) {
	files = make([]string, 0, 10)
	suffix = strings.ToUpper(suffix)

	err = filepath.Walk(dir_path, func(file_name string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, file_name)
		}
		return nil
	})

	return files, err
}


