// Author: xufei
// Date: 2019/4/30

package data

import (
	"fmt"
	"io"
	"os"
)

type DiskModel struct {
}

func (s *DiskModel) Open(name string) (io.ReadWriteCloser, error) {
	fmt.Println("disk open...")
	return os.Open(name)
}

func newDiskStorage() *DiskModel {
	return &DiskModel{}
}

type TempModel struct {
}

func (s *TempModel) Open(name string) (io.ReadWriteCloser, error) {
	fmt.Println("temp open...")
	return os.Open(name)
}

func newTempStorage() *TempModel {
	return &TempModel{}
}

type MemoryModel struct {
}

func (s *MemoryModel) Open(name string) (io.ReadWriteCloser, error) {
	fmt.Println("memory open...")
	return os.Open(name)
}

func newMemoryStorage() *MemoryModel {
	return &MemoryModel{}
}
