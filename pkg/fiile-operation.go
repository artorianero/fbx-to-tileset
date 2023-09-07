package tools

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func ReadFile(path string) []byte {
	data, err := ReadFileWithError(path)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return data
}

func ReadFileWithError(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		logrus.Error(err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	logrus.Info("filesize:", fi.Size())

	data := make([]byte, fi.Size())
	n, err := f.Read(data)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Infof("读取到%d个字符。", n)

	return data[:n], nil
}

func WriteFile(path string, data []byte) {
	err := WriteFileWithError(path, data)
	if err != nil {
		logrus.Error(err)
	}
}

func WriteFileWithError(path string, data []byte) error {
	dir := GetDirFromPath(path)
	if dir != "" {
		err := os.MkdirAll(dir, 0644)
		if err != nil {
			return err
		}
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	n, err := f.Write(data)
	if err != nil {
		return err
	}
	if n < 1 {
		return errors.New("write 0 bytes")
	}
	return nil
}

// 读取目录下的所有文件内容
func ReadDir(path string) (map[string][]byte, error) {
	m := make(map[string][]byte)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if !fi.IsDir() {
			fileName := strings.Replace(fi.Name(), "\\", "/", -1)
			fullPath := path + "/" + fileName
			data, err := ReadFileWithError(fullPath)
			if err != nil {
				return nil, err
			} else {
				m[fileName] = data
			}
		} else {
			continue
		}
	}
	return m, nil
}

// 读取目录下的所有文件名称（基于输入参数的全路径，仅返回当前目录下的文件路径，不遍历子目录）
func ReadFileNameFromDir(path string) []string {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		logrus.Info(err)
		return nil
	}

	if path[len(path)-1] != '/' {
		path += "/"
	}

	var result []string
	for _, fi := range dir {
		if !fi.IsDir() {
			result = append(result, path+fi.Name())
		}
	}

	return result
}

// 处理目录类型的路径，统一在最后添加"/"（如果原来有就维持原样）
func FixFolderPath(path string) string {
	if len(path) > 0 && path[len(path)-1] != '/' {
		path += "/"
	}
	return path
}

func GetDirFromPath(path string) string {
	dirs := strings.Split(path, "/")
	dir := ""
	for i := 0; i < len(dirs)-1; i++ {
		dir += dirs[i] + "/"
	}
	return dir
}
