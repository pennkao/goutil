package file

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
	if filename == "" {
		err = errors.New("not found in paths")
		return
	}
	for _, path := range paths {
		if fullPath = filepath.Join(path, filename); IsExist(fullPath) {
			return
		}
	}
	err = errors.New(fullPath + " not found in paths")
	return
}

// list dirs
func ReadDir(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := make([]string, 0, sz)
	for i := 0; i < sz; i++ {
		name := fs[i].Name()
		if name != "." && name != ".." {
			ret = append(ret, name)
		}
	}

	return ret, nil
}


// list dirs
func ReadDirDirs(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := make([]string, 0, sz)
	for i := 0; i < sz; i++ {
		if fs[i].IsDir() {
			name := fs[i].Name()
			if name != "." && name != ".." {
				ret = append(ret, name)
			}
		}
	}

	return ret, nil
}

// list files
func ReadDirFiles(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := make([]string, 0, sz)
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			ret = append(ret, fs[i].Name())
		}
	}

	return ret, nil
}
