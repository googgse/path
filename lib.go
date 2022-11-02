package path

import (
	// import local packages
	"github.com/googgse/path/lib"
)

func Exists(path string) bool {
	// return value
	return lib.Exists(path)
}

func IsDir(path string) bool {
	// return value
	return lib.IsDir(path)
}

func IsFile(path string) bool {
	// return value
	return lib.IsFile(path)
}
