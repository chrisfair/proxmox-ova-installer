package filesystem

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"bufio"

)

type ScannerInterface interface {
	Scan() bool
	Text() string
	Err() error
}

type ScannerWrapper struct {
	scanner *bufio.Scanner
}

func (sw *ScannerWrapper) Scan() bool {
	return sw.scanner.Scan()
}

func (sw *ScannerWrapper) Text() string {
	return sw.scanner.Text()
}

func (sw *ScannerWrapper) Err() error {
	return sw.scanner.Err()
}

// Interfaces for file operations and command execution, making them easier to mock.
type FileSystem interface {
	Open(name string) (io.ReadCloser, error)
	Create(name string) (io.WriteCloser, error)
	MkdirAll(path string, perm os.FileMode) error
	Remove(name string) error
	Walk(root string, walkFn filepath.WalkFunc) error
	Copy(dst io.Writer, src io.Reader) error
	GzNewReader(r io.Reader) (*gzip.Reader, error)
	TarNewReader(r io.Reader) *tar.Reader
	TarTypeDir() byte
	ReadFile(name string) ([]byte, error)
	WriteFile(name string, fileContents []byte, filePermissions os.FileMode) error
	Stat(name string) (os.FileInfo, error)
	JSONMarshal(v any) ([]byte, error)
	BufioNewScanner(r io.Reader) ScannerInterface
	Chmod(fileName string, fileMode os.FileMode) error
	RemoveAll(path string) error
}

// Default implementations of the interfaces.
type DefaultFileSystem struct{}

func (fs *DefaultFileSystem) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (fs *DefaultFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (fs *DefaultFileSystem) Chmod(fileName string, fileMode os.FileMode) error {
	return os.Chmod(fileName, fileMode) 
}

func (fs *DefaultFileSystem) WriteFile(
	name string, 
	fileContents []byte, 
	filePermissions os.FileMode) error {
	return os.WriteFile(name, fileContents, filePermissions)
}

func (fs *DefaultFileSystem) Open(name string) (io.ReadCloser, error) {
	return os.Open(name)
}

func (fs *DefaultFileSystem) Create(name string) (io.WriteCloser, error) {
	return os.Create(name)
}

func (fs *DefaultFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (fs *DefaultFileSystem) Remove(name string) error {
	return os.Remove(name)
}

func (fs *DefaultFileSystem) Walk(root string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(root, walkFn)
}

func (fs *DefaultFileSystem) Copy(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	return err
}

func (fs *DefaultFileSystem) GzNewReader(r io.Reader) (*gzip.Reader, error) {
	return gzip.NewReader(r)
}

func (fs *DefaultFileSystem) TarNewReader(r io.Reader) *tar.Reader {
	return tar.NewReader(r)
}

func (fs *DefaultFileSystem) TarTypeDir() byte {
	return tar.TypeDir
}

func (fs *DefaultFileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (fs *DefaultFileSystem) BufioNewScanner(r io.Reader) ScannerInterface{
	return &ScannerWrapper{
		scanner: bufio.NewScanner(r),
	}
}


type DefaultChecksumVerifier struct{
	fs FileSystem
}

func (dcv *DefaultChecksumVerifier) Verify(filePath, expectedChecksum string) error {
	file, err := dcv.fs.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	hash := sha256.New()
	if err := dcv.fs.Copy(hash, file); err != nil {
		return err
	}

	actualChecksum := fmt.Sprintf("%x", hash.Sum(nil))
	if actualChecksum != expectedChecksum {
		return fmt.Errorf("checksum mismatch for %s: expected %s, got %s", filePath, expectedChecksum, actualChecksum)
	}

	return nil
}

func (fs *DefaultFileSystem) JSONMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

