package filesys

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

type OsFsHandler struct {
	afero.Afero
}

func NewOsFsHandler() *OsFsHandler {
	return &OsFsHandler{
		Afero: afero.Afero{Fs: afero.NewOsFs()},
	}
}

func (fs *OsFsHandler) MkdirAllNoExist(dir string, perm fs.FileMode) error {
	ok, err := fs.DirExists(dir)
	if err != nil {
		return err
	}
	if !ok {
		err = fs.MkdirAll(dir, perm)
	}
	return err
}
func (fs *OsFsHandler) WriteDirsFile(src string, data []byte, perm fs.FileMode) error {
	err := fs.MkdirAllNoExist(filepath.Dir(src), perm)
	if err != nil {
		return err
	}
	err = fs.WriteFile(src, data, perm)
	return err
}

func (fs *OsFsHandler) CopyFile(src string, dst string, perm fs.FileMode, overwrite bool) error {
	srcR, err := fs.Open(src)
	if err != nil {
		return err
	}
	defer srcR.Close()
	if !overwrite {
		ok, err := fs.Exists(dst)
		if err != nil || ok {
			return err
		}
	}
	dstW, err := fs.Create(dst)
	if err != nil {
		return err
	}
	defer dstW.Close()
	_, err = io.Copy(dstW, srcR)
	if err != nil {
		return err
	}
	dstInfo, err := fs.Stat(dst)
	if err != nil || dstInfo.Mode().Perm() == perm {
		return err
	}
	err = fs.Chmod(dst, perm)
	return err
}

func (fs *OsFsHandler) CopyDir(src, dst string, perm fs.FileMode, overwrite bool) error {
	dirContents, err := fs.ReadDir(src)
	if err != nil {
		return err
	}

	err = fs.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	for _, item := range dirContents {
		srcFile := filepath.Join(src, item.Name())
		dstFile := filepath.Join(dst, item.Name())

		if item.IsDir() {
			err = fs.CopyDir(srcFile, dstFile, perm, overwrite)
		} else {
			err = fs.CopyFile(srcFile, dstFile, perm, overwrite)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// 在文件末尾追加内容，如果文件及不存在则创建
func (fs *OsFsHandler) AppendFile(src string, data []byte, perm fs.FileMode) error {
	err := fs.MkdirAllNoExist(filepath.Dir(src), perm)
	if err != nil {
		return err
	}
	f, err := fs.OpenFile(src, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

// 在文件头部追加内容，如果文件及不存在则创建
func (fs *OsFsHandler) PrependFile(src string, data []byte, perm fs.FileMode) error {
	ok, err := fs.Exists(src)
	if err != nil {
		return err
	}
	if !ok {
		return fs.WriteFile(src, data, perm)
	}

	f, err := fs.OpenFile(src, os.O_RDWR, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	old, err := fs.ReadFile(src)
	if err != nil {
		return nil
	}
	data = append(data, old...)
	fs.WriteFile(src, data, perm)
	return err
}
