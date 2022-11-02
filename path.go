package path

import (
	// import built-in packages
	"os"
	"path"
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	// control flows
	if err != nil {
		// return value
		return os.IsExist(err)
	}
	// return value
	return true
}

// 判断所给路径是否为绝对路径
func IsAbs(_path string) bool {
	// return value
	return path.IsAbs(_path)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	// control flows
	if err != nil {
		// return value
		return false
	}
	// return value
	return stat.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	// return value
	return !IsDir(path)
}

// 将任意数量的路径元素放入一个单一路径
func Join(_path ...string) string {
	return path.Join(_path...)
}
