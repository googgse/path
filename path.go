package path

import (
	// import local packages
	"github.com/googgse/path/lib"
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	// return value
	return lib.Exists(path)
}

// 判断所给路径是否为绝对路径
func IsAbs(path string) bool {
	// return value
	return lib.IsAbs(path)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	// return value
	return lib.IsDir(path)
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	// return value
	return lib.IsFile(path)
}
