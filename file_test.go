package helper

import (
	"os"
	"testing"
)

func TestIsExist(t *testing.T) {
	filename := "./file.go"
	if !TValidate.IsFileExist(filename) {
		t.Error("file does not exist")
		return
	}
}

func TestWriteFile(t *testing.T) {
	str := []byte("Hello World!")
	err := TFile.FileWrite("./temp/file", str)
	if err != nil {
		t.Error("file get contents unit test fail")
		return
	}
	_ = TFile.FileWrite("/root/hello/world", str)
}

func TestGetExt(t *testing.T) {
	filename := "./file.go"
	if TFile.GetExt(filename) != "go" {
		t.Error("file extension unit test fail")
		return
	}

	TFile.GetExt("./temp/file")
}

func TestBasename(t *testing.T) {
	path := "./temp/go.png"
	res := TFile.Basename(path)
	if res != "go.png" {
		t.Error("Basename unit test fail")
		return
	}
}

func TestGetMime(t *testing.T) {
	filename := "./temp/go.png"
	mime1 := TFile.GetMime(filename, true)
	mime2 := TFile.GetMime(filename, false)

	if mime1 != mime2 {
		t.Error("GetMime unit test fail")
		return
	}

	TFile.GetMime("./temp/go-lnk", false)
	TFile.GetMime("./", false)
}

func TestReadInArray(t *testing.T) {
	filepath := "./temp/go.txt"
	arr, err := TFile.ReadInArray(filepath)
	if err != nil || len(arr) != 1 {
		t.Error("ReadInArray unit test fail")
		return
	}

	_, _ = TFile.ReadInArray("./hello")
}

func TestReadFile(t *testing.T) {
	filename := "./file.go"
	cont, _ := TFile.ReadFile(filename)
	if string(cont) == "" {
		t.Error("file get contents unit test fail")
		return
	}
}

func TestFileSize(t *testing.T) {
	filename := "./file.go"
	if TFile.FileSize(filename) <= 0 {
		t.Error("file size unit test fail")
		return
	}
	TFile.FileSize("./hello")
}

func TestTouchRenameRemove(t *testing.T) {
	file1 := "./temp/test/zero"
	file2 := "./temp/test/2m"
	file3 := "./temp/test/"
	file4 := "./temp/test/"

	//创建文件
	res1 := TFile.Touch(file1, 0)
	res2 := TFile.Touch(file2, 2097152)
	if !res1 || !res2 {
		t.Error("Touch unit test fail")
		return
	}

	//重命名
	file5 := "./temp/test/zero_re"
	file6 := "./temp/test/2m_re"
	err1 := TFile.Rename(file1, file5)
	err2 := TFile.Rename(file2, file6)
	if err1 != nil || err2 != nil {
		t.Error("Unlink unit test fail")
		return
	}

	//删除文件
	err3 := TFile.Remove(file5)
	err4 := TFile.Remove(file6)
	if err3 != nil || err4 != nil {
		t.Error("Unlink unit test fail")
		return
	}

	TFile.Touch(file3, 0)
	TFile.Touch(file4, 0)
}

func TestDelDir(t *testing.T) {
	dir := "./temp"
	err := TFile.DelDir(dir, true)
	if err != nil || TValidate.IsDir(dir) {
		t.Error("DelDir unit test fail")
		return
	}

	_ = TFile.DelDir("./hello", true)
	_ = TFile.DelDir("/root", true)

}

func TestAbsPath(t *testing.T) {
	filename := "./temp/go.png"
	absPath := TFile.AbsPath(filename)
	if !TValidate.IsFileExist(absPath) {
		t.Error("file does not exist")
		return
	}
	TFile.AbsPath("")
	TFile.AbsPath("file:///c:/test.go")

	//手工引发 filepath.Abs 错误
	//创建目录
	testPath := "./temp/abs/123"
	err := os.MkdirAll(testPath, 0755)
	if err == nil {
		//当前目录
		testDir, _ := os.Getwd()

		filename = "../../go.png"
		//进入目录
		_ = os.Chdir(testPath)
		pwdir, _ := os.Getwd()

		//删除目录
		_ = os.Remove(pwdir)

		//再获取路径
		res := TFile.AbsPath(filename)

		if res != "D:\\php\\helper\\temp\\go.png" {
			t.Error("KFile.AbsPath unit test fail")
		}

		//回到旧目录
		_ = os.Chdir(testDir)
	}
}
