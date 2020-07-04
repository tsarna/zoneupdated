package atomicfile

import (
  "fmt"
  "os"
)

type AtomicFile struct {
  fileName     string
  tempFileName string
  tempFile     *os.File
}

const (
  tempSuffix = ".tmp"
)

func Open(filename string) (*AtomicFile, error) {
  tempFileName := fmt.Sprint(filename, tempSuffix)
  tempFile, err := os.Create(tempFileName)
  if (err != nil) {
    return nil, err
  }

  return &AtomicFile{filename, tempFileName, tempFile}, nil
}

func (a *AtomicFile) Write(p []byte) (n int, err error) {
  if (a.tempFile == nil) {
    return 0, fmt.Errorf("File is closed")
  }

  return a.tempFile.Write(p)
}

func (a *AtomicFile) Abort() error {
  a.close()
  return os.Remove(a.tempFileName)
}

func (a *AtomicFile) Commit() error {
  a.close()
  return os.Rename(a.tempFileName, a.fileName)
}

func (a *AtomicFile) close() {
  a.tempFile.Close()
  a.tempFile = nil
}