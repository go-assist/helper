package helper

import (
	"errors"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// IsFileExist 文件是否存在
func (tf *TsFile) IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

// FileWrite 写文件
func (tf *TsFile) FileWrite(filePath string, data []byte) (err error) {

	if dir := path.Dir(filePath); dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filePath, data, 0655)
}

// GetExt 获取文件扩展名
func (tf *TsFile) GetExt(filePath string) string {
	suffix := filepath.Ext(filePath)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}
	return suffix
}

// Basename 返回路径中的文件名部分
func (tf *TsFile) Basename(filePath string) string {
	return filepath.Base(filePath)
}

// GetMime 获取mime类型
// fast为true时根据后缀快速获取;为false时读取文件头获取.
func (tf *TsFile) GetMime(filePath string, fast bool) string {
	var res string
	if fast {
		suffix := filepath.Ext(filePath)
		res = mime.TypeByExtension(suffix)
	} else {
		srcFile, err := os.Open(filePath)
		if err != nil {
			return res
		}

		buffer := make([]byte, 512)
		_, err = srcFile.Read(buffer)
		if err != nil {
			return res
		}

		res = http.DetectContentType(buffer)
	}

	return res
}

// ReadInArray 文件写入数组
func (tf *TsFile) ReadInArray(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

// ReadFile 读取文件内容
func (tf *TsFile) ReadFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	return data, err
}

// FileSize 获取文件大小
func (tf *TsFile) FileSize(filePath string) int64 {
	f, err := os.Stat(filePath)
	if nil != err {
		return -1
	}
	return f.Size()
}

// Rename 重命名
func (tf *TsFile) Rename(oldName string, newName string) error {
	return os.Rename(oldName, newName)
}

// Touch 快速创建指定大小的文件,size为字节
func (tf *TsFile) Touch(filePath string, size int64) bool {
	//创建目录
	destDir := filepath.Dir(filePath)
	if destDir != "" && !tf.IsDir(destDir) {
		if err := os.MkdirAll(destDir, 0766); err != nil {
			return false
		}
	}

	fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return false
	}
	defer func(fd *os.File) {
		err = fd.Close()
		if err != nil {
			return
		}
	}(fd)

	if size > 1 {
		_, _ = fd.Seek(size-1, 0)
		_, _ = fd.Write([]byte{0})
	}
	return true
}

// Remove 删除文件
func (tf *TsFile) Remove(filePath string) error {
	return os.Remove(filePath)
}

// DelDir 删除目录;delRoot为true时连该目录一起删除;为false时只清空该目录.
func (tf *TsFile) DelDir(dir string, delRoot bool) error {
	realPath := tf.AbsPath(dir)
	if !tf.IsDir(realPath) {
		return errors.New("Dir does not exists:" + dir)
	}

	names, err := ioutil.ReadDir(realPath)
	if err != nil {
		return err
	}

	for _, retry := range names {
		file := path.Join([]string{realPath, retry.Name()}...)
		err = os.RemoveAll(file)
	}
	//删除根节点(指定的目录)
	if delRoot {
		err = os.RemoveAll(realPath)
	}
	return err
}

// AbsPath 获取绝对路径,path可允许不存在.
func (tf *TsFile) AbsPath(filePath string) string {
	fullPath := ""
	res, err := filepath.Abs(filePath) // filepath.Abs最终使用到os.GetWorkList()检查
	if err != nil {
		fullPath = filepath.Clean(filepath.Join(`/`, filePath))
	} else {
		fullPath = res
	}

	return fullPath
}

// IsDir 是否目录(且存在)
func (tf *TsFile) IsDir(filePath string) bool {
	f, err := os.Lstat(filePath)
	if os.IsNotExist(err) || nil != err {
		return false
	}
	return f.IsDir()
}
