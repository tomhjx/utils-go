package file

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getAppRootPath() string {
	pwd, _ := os.Getwd()
	rootPath := filepath.Dir(filepath.Dir(pwd))
	return rootPath
}
func TestOsFsWriteFile(t *testing.T) {
	rootPath := getAppRootPath()
	hd := NewOsFsHandler()
	c := []byte("TestOsFsWriteFile," + time.Now().Format("2006-01-02 15:04:05"))
	src := rootPath + "/tmp/utils-go-test.txt"
	err := hd.WriteFile(src, c, 0777)
	assert.Nil(t, err)
}

func TestOsFsWriteDirsFile(t *testing.T) {
	rootPath := getAppRootPath()
	hd := NewOsFsHandler()
	c := []byte("TestOsFsWriteDirsFile," + time.Now().Format("2006-01-02 15:04:05"))
	src := rootPath + "/tmp/a/b/c/utils-go-test.txt"
	err := hd.WriteDirsFile(src, c, 0777)
	assert.Nil(t, err)
}

func TestOsFsCopyFile(t *testing.T) {
	rootPath := getAppRootPath()
	hd := NewOsFsHandler()
	c := []byte("TestOsFsWriteFile," + time.Now().Format("2006-01-02 15:04:05"))
	src := rootPath + "/tmp/utils-go-test.txt"
	err := hd.WriteFile(src, c, 0777)
	assert.Nil(t, err)
	dest := rootPath + "/tmp/utils-go-test-copy.txt"
	err = hd.CopyFile(src, dest, 0777, true)
	assert.Nil(t, err)

	dest = rootPath + "/tmp/utils-go-test-copy-no-ow.txt"
	err = hd.CopyFile(src, dest, 0777, false)
	assert.Nil(t, err)

}

func TestOsFsCopyDir(t *testing.T) {
	rootPath := getAppRootPath()
	hd := NewOsFsHandler()
	c := []byte("TestOsFsWriteDirsFile," + time.Now().Format("2006-01-02 15:04:05"))
	src := rootPath + "/tmp/a/b/c/utils-go-test.txt"
	err := hd.WriteDirsFile(src, c, 0777)
	assert.Nil(t, err)

	src = rootPath + "/tmp/a/b/cc/utils-go-test.txt"
	err = hd.WriteDirsFile(src, c, 0777)
	assert.Nil(t, err)

	src = rootPath + "/tmp/a"
	dst := rootPath + "/tmp/a-copy"
	err = hd.CopyDir(src, dst, 0777, true)
	assert.Nil(t, err)

	src = rootPath + "/tmp/a"
	dst = rootPath + "/tmp/a-copy-no-ow"
	err = hd.CopyDir(src, dst, 0777, false)
	assert.Nil(t, err)

}
func TestOsFsAppendFile(t *testing.T) {
	rootPath := getAppRootPath()
	hd := NewOsFsHandler()
	c := []byte("TestOsFsAppendFile," + time.Now().Format("2006-01-02 15:04:05\r\n"))
	src := rootPath + "/tmp/utils-go-test-append.txt"
	err := hd.AppendFile(src, c, 0777)
	assert.Nil(t, err)
	src = rootPath + "/tmp/aa/utils-go-test-append.txt"
	err = hd.AppendFile(src, c, 0777)
	assert.Nil(t, err)
}

func TestOsFsPrependFile(t *testing.T) {
	rootPath := getAppRootPath()
	hd := NewOsFsHandler()
	c := []byte("TestOsFsPrependFile," + time.Now().Format("2006-01-02 15:04:05\r\n"))
	src := rootPath + "/tmp/utils-go-test-prepend.txt"
	err := hd.PrependFile(src, c, 0777)
	assert.Nil(t, err)
	src = rootPath + "/tmp/aa/utils-go-test-prepend.txt"
	err = hd.PrependFile(src, c, 0777)
	assert.Nil(t, err)
}
