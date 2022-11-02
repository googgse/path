package path

import (
	// import built-in packages
	"os"
	"path"
)

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

func IsAbs(_path string) bool {
	// return value
	return path.IsAbs(_path)
}

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

func IsFile(path string) bool {
	// return value
	return !IsDir(path)
}
