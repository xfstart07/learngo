// Author: Xu Fei
// Date: 2018/9/13
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
)

const filechunk = 8192 // we settle for 8KB

func main() {
	filePath := "/Users/x/Downloads/201809111536630287271_鲸唱扫码2.mkv"

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// calculate the file size
	info, _ := file.Stat()

	filesize := info.Size()

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	fmt.Printf("%s checksum is %x\n", file.Name(), hash.Sum(nil))
}
